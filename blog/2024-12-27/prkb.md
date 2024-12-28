# PRKB

start here:

~~~
sec2\src\pr\v0404\trustedexec\base\drmteexbformat_generated.h
#define DRM_TEE_XB_KB_FORMAT_ID        XB_DEFINE_DWORD_FORMAT_ID( 'P', 'R', 'K', 'B' )
~~~

`DRM_TEE_XB_KB_FORMAT_ID` reference:

~~~
sec2\src\pr\v0404\trustedexec\base\drmteexbformat_generated.c
static DRM_GLOBAL_CONST
DRM_XB_HEADER_DESCRIPTION
s_DRM_TEE_XB_KB_HeaderDescription_DRM_PRKB_CURRENT_VERSION
PR_ATTR_DATA_OVLY(_s_DRM_TEE_XB_KB_HeaderDescription_DRM_PRKB_CURRENT_VERSION) = {
   XB_FORMAT_ID_LENGTH_DWORD, DRM_TEE_XB_KB_FORMAT_ID,
   DRM_PRKB_CURRENT_VERSION, XB_ALIGN_8_BYTE, 0, NULL
};
~~~

`DRM_GLOBAL_CONST` definition:

~~~
sec2\src\pr\v0404\inc\drmcompiler.h
#define DRM_GLOBAL_CONST const
~~~

`DRM_XB_HEADER_DESCRIPTION` definition:

~~~
sec2\src\pr\v0404\inc\drmxb.h
typedef struct __tagDRM_XB_HEADER_DESCRIPTION
{
   XB_FORMAT_ID_LENGTH             eFormatIdLength;
   
   /* 4 or 8 byte string identifier.  Use XB_DEFINE_DWORD_FORMAT_ID or XB_DEFINE_QWORD_FORMAT_ID to create this value */
   DRM_UINT64                      qwFormatIdentifier;
   
   /* Version supported by this description */
   DRM_DWORD                       dwFormatVersion;  
   
   /* Byte alignment identifier.  Must be a non-zero, 4^x value (e.g. 1, 4, 8, 16, etc). */
   DRM_XB_ALIGNMENT                eAlign;  
   
   /* Where in the containing structure should the data be serialized from/to */
   DRM_WORD                        wOffsetOfSizeInHeaderStruct;   
   
   /* The header can have optional extra data -- but only 1 */
   const DRM_XB_ENTRY_DESCRIPTION *pEntryDescription;     
   
} DRM_XB_HEADER_DESCRIPTION;
~~~
