#include <stdio.h>
#include <gdal.h>
#include <gdal_utils.h>
#include "goodle.h"

char *Info(char *fname)
{
  GDALAllRegister();
  char *pgArg = "gdal2";
  char *jsonFlag = "-json";
  GDALDatasetH hDataSet = GDALOpenEx(fname, GDAL_OF_READONLY | GDAL_OF_RASTER | GDAL_OF_VERBOSE_ERROR, CPL_NULLPTR, CPL_NULLPTR, CPL_NULLPTR);
  char *arg[3] = {pgArg, jsonFlag};
  GDALInfoOptions *infoOptions = GDALInfoOptionsNew(arg, CPL_NULLPTR);
  return GDALInfo(hDataSet, infoOptions);
}

