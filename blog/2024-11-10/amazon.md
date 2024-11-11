# amazon

~~~
com\amazon\identity\auth\device\C4655v5.java
private static final String[] f5928a = {"-----BEGIN RSA PRIVATE KEY-----",
"-----END RSA PRIVATE KEY-----", "-----BEGIN EC PRIVATE KEY-----", "-----END EC
PRIVATE KEY-----", "-----BEGIN PRIVATE KEY-----", "-----END PRIVATE KEY-----"};
~~~

then:

~~~java
if (str != null) {
   String[] strArr = f5928a;
   for (int i2 = 0; i2 < 6; i2++) {
       str = str.replace(strArr[i2], "");
   }
   str2 = str.trim();
}
~~~

then:

~~~java
public static PrivateKey m2834b(String str) throws InvalidKeySpecException
~~~

then:

~~~
AbstractC4493m5.java
29:                    this.f5297b = C4655v5.m2834b(mo927a().m1316a());
~~~

then:

~~~
C4293b.java
public final String m1316a() {
  return this.f4560b;
}
~~~

then:

~~~java
public C4293b(String str, String str2) {
  this.f4559a = str;
  this.f4560b = str2;
}
~~~

then:

~~~
C4437i1.java
c4293b = new C4293b((String)
mo2419a2.get(AccountConstants.TOKEN_TYPE_DEVICE_ADP_TOKEN), (String)
mo2419a2.get(AccountConstants.TOKEN_TYPE_DEVICE_PRIVATE_KEY));
~~~

then:

~~~
utils\AccountConstants.java
public static final String TOKEN_TYPE_DEVICE_PRIVATE_KEY =
"com.amazon.dcp.sso.token.device.privatekey";
~~~

then:

~~~
com\amazon\identity\auth\device\C4437i1.java
HashMap mo2419a2 =
abstractC4585f.mo2419a("com.amazon.identity.auth.device.credentials.AnonymousAccountCredentials.SHARED_PREFS",
list);
~~~

then:

~~~
com\amazon\identity\auth\device\C4437i1.java
public static C4293b m1878a(AbstractC4585f abstractC4585f) {
~~~

then:

~~~
com\amazon\identity\auth\device\C4437i1.java
public final C4293b mo927a() {
  C4293b c4293b = this.f5111f;
  if (c4293b == null || TextUtils.isEmpty(c4293b.m1316a()) || TextUtils.isEmpty(this.f5111f.m1317b())) {
      this.f5111f = m1878a(this.f5110e);
  }
  return this.f5111f;
}
~~~

then:

~~~java
com\amazon\identity\auth\device\C4437i1.java
this.f5110e = (
   (InterfaceC4239a4) m2953a.getSystemService("dcp_data_storage_factory")
).mo1183a();
~~~

then:

~~~java
com\amazon\identity\auth\device\C4696y9.java
} else if ("dcp_data_storage_factory".equals(str)) {
    obj = C4586g.m2425a(this);
~~~
