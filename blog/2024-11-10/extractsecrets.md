# extractsecrets

~~~
msprcp> extractsecrets mspr_binary\libstd_cai_client_drm_msplayready.so
~~~

<https://security-explorations.com/materials/mspr_toolkit_README.md.txt>

fail:

https://apkmirror.com/apk/bell-media-inc/ctv

fail:

https://apkmirror.com/apk/bell-media-inc/ctv-android-tv

fail:

~~~
> play -i be.rtbf.auvio
details[8] = 0 USD
details[13][1][4] = 3.1.42
details[13][1][16] = Jun 26, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 8.0 and up
details[15][18] = https://www.rtbf.be/charte/detail
downloads = 1.77 million
name = RTBF Auvio : direct et replay
size = 48.09 megabyte
version code = 1301042
~~~

fail:

~~~
> play -i com.amcplus.amcfullepisodes
details[8] = 0 USD
details[13][1][4] = 1.8.21.2
details[13][1][16] = Oct 23, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = https://www.amcnetworks.com/privacy-policy/
downloads = 1.66 million
name = AMC+
size = 45.00 megabyte
version code = 172970545
~~~

---------------------------------------------------------------------------------

## amazon

~~~
> play -i com.amazon.avod.thirdpartyclient -abi armeabi-v7a
details[8] = 0 USD
details[13][1][4] = 3.0.388.545
details[13][1][16] = Nov 8, 2024
details[13][1][17] = APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = http://www.amazon.com/gp/help/customer/display.html?nodeId=468496
downloads = 692.53 million
name = Amazon Prime Video
size = 51.13 megabyte
version code = 388000545
~~~

works:

~~~
com.amazon.avod.thirdpartyclient-388000545> rg -a CHAI\b
lib\armeabi-v7a\libAIVPlayReadyLicensing.so
�r5�&wv���|(ʖDG�gF37B��ʾ���l�]�R�7yޙ������?��%o�.@@�J      �����S/Eι�d�U���f���\�`&BR�~ױ�@8���+?�^lǰ�=g�y�3
. h|.f}'w�@ޠW�`�Lm嗲�{�        y������q7     ��<��O����R�o�'K`��xQ~h��� ����r�_�|��8ee����:U��+��~�(8j���S3�;}��>��B&�)�V�����6�0��N�U�"��X7CHAI\CERT�,X�LA������4���/��Ч?eۦ�oF�4�EjTO��,�G�:O�y�������`5,��������_��^�_��(��3�
~~~

## cinemember

fail:

https://apkpure.com/cinemember/nl.peoplesplayground.audienceplayer.cinemember

## criterionchannel

fail:

~~~
> play -i com.criterionchannel
details[8] = 0 USD
details[13][1][4] = 8.806.1
details[13][1][16] = Oct 23, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = https://www.criterionchannel.com/privacy
downloads = 210.95 thousand
name = The Criterion Channel
size = 67.33 megabyte
version code = 11280
~~~

## draken

fail:

https://apkcombo.com/draken-film/com.draken.android

## hulu

fail:

~~~
> play -i com.hulu.plus
details[8] = 0 USD
details[13][1][4] = 5.9.1+15066-google
details[13][1][16] = Oct 25, 2024
details[13][1][17] = APK
details[13][1][82][1][1] = 7.1 and up
details[15][18] = https://www.hulu.com/privacy
downloads = 95.44 million
name = Hulu: Stream TV shows & movies
size = 19.62 megabyte
version code = 5015066
~~~

next:

~~~
> play -i com.hulu.livingroomplus -abi armeabi-v7a -leanback
details[8] = 0 USD
details[13][1][4] = 5EEC9D81P3.9.846
details[13][1][16] = Aug 6, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 4.4 and up
details[15][18] = https://www.hulu.com/privacy
downloads = 20.39 million
name = Hulu for Android TV
size = 47.67 megabyte
version code = 3009846
~~~

pass:

~~~
> rg -a CHAI\b
com.hulu.livingroomplus-config.armeabi_v7a-3009846\lib\armeabi-v7a\libwkf_support.so
163933:CHAI<CERT�
�U� ��?��^P�����N7䱣�k˱�/����l�RG�{��S4Hulu LLCWiiUWiiU�@`��ϡ-��s(�*��f���{��r!�#���g�DO��L��G��8��6�J�t���J���C��Lؓ��Lh#��C�ͪ�
�� �C�[��W'�o��YQy��h`M�X��,��
~~~

## itv

fail:

https://apkmirror.com/apk/itv-plc/itv-hub

## max

fail:

~~~
> play -i com.wbd.stream
details[8] = 0 USD
details[13][1][4] = 5.0.0.67
details[13][1][16] = Oct 23, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = https://www.max.com/privacy
downloads = 83.80 million
name = Max: Stream HBO, TV, & Movies
size = 114.30 megabyte
version code = 35354161
~~~

## mubi

fail:

~~~
> play -i com.mubi
details[8] = 0 USD
details[13][1][4] = 58.0
details[13][1][16] = Oct 28, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 5.1 and up
details[15][18] = https://mubi.com/privacy_policy
downloads = 5.51 million
name = MUBI: Curated Cinema
size = 26.12 megabyte
version code = 29145342
~~~

## nbc

fail:

~~~
>  play -i com.nbcuni.nbc
details[8] = 0 USD
details[13][1][4] = 9.13.1
details[13][1][16] = Oct 16, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = https://www.nbcuniversal.com/privacy?intake=NBC_Entertainment
downloads = 16.54 million
name = The NBC App - Stream TV Shows
size = 73.71 megabyte
version code = 2000005553
~~~

## paramount

fail:

~~~
> play -i com.cbs.app
details[8] = 0 USD
details[13][1][4] = 15.0.46
details[13][1][16] = Oct 30, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = http://legalterms.cbsinteractive.com/privacy
downloads = 36.54 million
name = Paramount+
size = 89.49 megabyte
version code = 211504644
~~~

## plex

fail:

~~~
> play -i com.plexapp.android
details[8] = 0 USD
details[13][1][4] = 10.23.0.1627
details[13][1][16] = Oct 21, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 6.0 and up
details[15][18] = https://plex.tv/legal
downloads = 66.85 million
name = Plex: Stream Movies & TV
size = 90.43 megabyte
version code = 963667377
~~~

## pluto

fail:

~~~
> play -i tv.pluto.android
details[8] = 0 USD
details[13][1][4] = 5.49.0
details[13][1][16] = Oct 29, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = https://corporate.pluto.tv/privacy-policy/
downloads = 220.01 million
name = Pluto TV: Watch Free Movies/TV
size = 55.33 megabyte
version code = 410400212
~~~

## rakuten

fail:

https://apkmirror.com/apk/rakuten-tv/rakuten-tv-movies-tv-series-android-tv

## roku

fail:

~~~
> play -i com.roku.web.trc -leanback
details[8] = 0 USD
details[13][1][4] = 1.1.19
details[13][1][16] = Oct 21, 2024
details[13][1][17] = APK APK APK APK
details[13][1][82][1][1] = 5.0 and up
details[15][18] = https://docs.roku.com/published/userprivacypolicy/en/us
downloads = 750.19 thousand
name = The Roku Channel
size = 9.94 megabyte
version code = 10019
~~~

## tubitv

fail:

~~~
> play -i com.tubitv -abi armeabi-v7a
details[8] = 0 USD
details[13][1][4] = 8.22.0
details[13][1][16] = Nov 4, 2024
details[13][1][17] = APK APK APK
details[13][1][82][1][1] = 9 and up
details[15][18] = https://tubitv.com/static/privacy
downloads = 173.99 million
name = Tubi: Free Movies & Live TV
size = 75.82 megabyte
version code = 851
~~~
