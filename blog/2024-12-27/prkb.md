# PRKB

start here:

~~~
sec2\src\pr\v0404\trustedexec\base\drmteexbformat_generated.h
#define DRM_TEE_XB_KB_FORMAT_ID        XB_DEFINE_DWORD_FORMAT_ID( 'P', 'R', 'K', 'B' )
~~~

`DRM_TEE_XB_KB_FORMAT_ID` reference:

~~~
sec2\src\pr\v0404\trustedexec\base\drmteexbformat_generated.c
{ XB_FORMAT_ID_LENGTH_DWORD, DRM_TEE_XB_KB_FORMAT_ID, DRM_PRKB_CURRENT_VERSION, XB_ALIGN_8_BYTE, 0, NULL};
~~~
