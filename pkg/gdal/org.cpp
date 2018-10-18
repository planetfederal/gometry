/*
MIT License

Copyright (c) 2018 Boundless

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

#include "ogr.h"
#include <set>
#include <gdal.h>
#include <ogr_api.h>
#include <ogr_srs_api.h>
#include <json-c/json.h>

void ReportOnLayer(OGRLayerH poLayer, json_object *json_output)
{
  const char *geomType = OGRGeometryTypeToName(OGR_L_GetGeomType(poLayer));
  json_object_object_add(json_output, "geometry", json_object_new_string(geomType));

  //Feature Count
  json_object_object_add(json_output, "feature_count", json_object_new_int64(OGR_L_GetFeatureCount(poLayer, 0)));

  //Extent
  json_object *extentJsonArray = json_object_new_array();
  OGREnvelope oExt;
  if (OGR_L_GetExtent(poLayer, &oExt, TRUE) == OGRERR_NONE)
  {
    json_object_array_add(extentJsonArray, json_object_new_double(oExt.MinX));
    json_object_array_add(extentJsonArray, json_object_new_double(oExt.MinY));
    json_object_array_add(extentJsonArray, json_object_new_double(oExt.MaxX));
    json_object_array_add(extentJsonArray, json_object_new_double(oExt.MaxY));
  }
  json_object_object_add(json_output, "extent", extentJsonArray);

  //Srs Wkt
  char *pszWkt = nullptr;
  if (OGR_L_GetSpatialRef(poLayer) != nullptr)
  {
    OGRSpatialReferenceH srs = OGR_L_GetSpatialRef(poLayer);
    OSRExportToWkt(srs, &pszWkt);
    json_object_object_add(json_output, "layer_srs_wkt", json_object_new_string(pszWkt));
  }

  //FID COlumn
  json_object_object_add(json_output, "fid_column", json_object_new_string(OGR_L_GetFIDColumn(poLayer)));
  json_object *geom_column_arr = json_object_new_array();
  OGRFeatureDefnH fdefh = OGR_L_GetLayerDefn(poLayer);
  OGRGeomFieldDefnH gfdh = OGR_FD_GetGeomFieldDefn(fdefh, 0);
  const char *nameRef = OGR_GFld_GetNameRef(gfdh);
  json_object_array_add(geom_column_arr, json_object_new_string(nameRef));
  json_object_object_add(json_output, "geometry_column", geom_column_arr);

  //Fields
  json_object *fieldsObject = json_object_new_object();
  OGRFeatureDefnH poDefn = OGR_L_GetLayerDefn(poLayer);
  for (int iAttr = 0; iAttr < OGR_FD_GetFieldCount(poDefn); iAttr++)
  {
    OGRFieldDefnH poField = OGR_FD_GetFieldDefn(poDefn, iAttr);
    json_object *jsonFieldDef = json_object_new_object();
    OGRFieldType subType = OGR_Fld_GetType(poField);
    const char *typeStr = OGR_GetFieldTypeName(subType);
    json_object_object_add(jsonFieldDef, "type", json_object_new_string(typeStr));
    json_object_object_add(jsonFieldDef, "width", json_object_new_int(OGR_Fld_GetWidth(poField)));
    json_object_object_add(jsonFieldDef, "precision", json_object_new_int(OGR_Fld_GetPrecision(poField)));
    json_object_object_add(fieldsObject, OGR_Fld_GetNameRef(poField), jsonFieldDef);
  }

  json_object_object_add(json_output, "fields", fieldsObject);
}

const char *OGRInfoJSONString(char *pszDataSource, char *layerName)
{
  OGRRegisterAll();
  const char *pgArg = "ogrinfo";
  const char *summ = "-summary";
  const char *arg[4] = {pgArg, summ, pszDataSource, layerName};
  GDALDatasetH poDS = GDALOpenEx(pszDataSource,
                                 GDAL_OF_VECTOR,
                                 CPL_NULLPTR, CPL_NULLPTR, CPL_NULLPTR);

  GDALDriverH poDriver = CPL_NULLPTR;
  if (poDS != nullptr)
  {
    poDriver = GDALGetDatasetDriver(poDS);
  }
  bool bVerbose = true;
  const char *desc = GDALGetDescription(poDriver);

  printf("INFO: Open of `%s'\n"
         "      using driver `%s' successful.\n",
         pszDataSource, desc);

  json_object *poJsonObject = json_object_new_object();
  json_object_object_add(poJsonObject, "driver", json_object_new_string(desc));
  json_object_object_add(poJsonObject, "file", json_object_new_string(pszDataSource));
  json_object_object_add(poJsonObject, "layer_name", json_object_new_string(layerName));

  json_object_object_add(poJsonObject, "geometry", json_object_new_string(layerName));

  std::set<OGRLayerH *> oSetLayers;
  while (true)
  {
    OGRLayerH poLayer = nullptr;
    OGRFeatureH poFeature = GDALDatasetGetNextFeature(poDS, &poLayer, nullptr, nullptr, nullptr);

    if (poLayer == nullptr)
    {
      break;
    }
    ReportOnLayer(poLayer, poJsonObject);
  }

  return json_object_to_json_string_ext(poJsonObject, JSON_C_TO_STRING_PRETTY);
}
