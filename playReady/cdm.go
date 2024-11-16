package playReady

import "encoding/xml"

type soap_envelope struct {
   XMLName xml.Name `xml:"soap:Envelope"`
   Soap    string   `xml:"xmlns:soap,attr"`
   Body    struct {
      AcquireLicense struct {
         Xmlns     string `xml:"xmlns,attr"`
         Challenge struct {
            Challenge struct {
               Xmlns string `xml:"xmlns,attr"`
               La    struct {
                  Xmlns         string `xml:"xmlns,attr"`
                  Id            string `xml:"Id,attr"`
                  Space         string `xml:"space,attr"`
                  Version       string `xml:"Version"`
                  ContentHeader struct {
                     WrmHeader struct {
                        Xmlns   string `xml:"xmlns,attr"`
                        Version string `xml:"version,attr"`
                        Data    struct {
                           ProtectInfo struct {
                              KeyLen string `xml:"keylen"`
                              AlgId  string `xml:"algid"`
                           } `xml:"PROTECTINFO"`
                           Kid string `xml:"kid"`
                        } `xml:"DATA"`
                     } `xml:"WRMHEADER"`
                  } `xml:"ContentHeader"`
                  Clientinfo struct {
                     ClientVersion string `xml:"CLIENTVERSION"`
                  } `xml:"CLIENTINFO"`
                  RevocationLists struct {
                     RevListInfo []struct {
                        ListId  string `xml:"ListID"`
                        Version string `xml:"Version"`
                     } `xml:"RevListInfo"`
                  } `xml:"RevocationLists"`
                  LicenseNonce  string `xml:"LicenseNonce"`
                  ClientTime    string `xml:"ClientTime"`
                  EncryptedData struct {
                     Xmlns            string `xml:"xmlns,attr"`
                     Type             string `xml:"Type,attr"`
                     EncryptionMethod struct {
                        Algorithm string `xml:"Algorithm,attr"`
                     } `xml:"EncryptionMethod"`
                     KeyInfo struct {
                        Xmlns        string `xml:"xmlns,attr"`
                        EncryptedKey struct {
                           Xmlns            string `xml:"xmlns,attr"`
                           EncryptionMethod struct {
                              Algorithm string `xml:"Algorithm,attr"`
                           } `xml:"EncryptionMethod"`
                           KeyInfo struct {
                              Xmlns   string `xml:"xmlns,attr"`
                              KeyName string `xml:"KeyName"`
                           } `xml:"KeyInfo"`
                           CipherData struct {
                              CipherValue string `xml:"CipherValue"`
                           } `xml:"CipherData"`
                        } `xml:"EncryptedKey"`
                     } `xml:"KeyInfo"`
                     CipherData struct {
                        CipherValue string `xml:"CipherValue"`
                     } `xml:"CipherData"`
                  } `xml:"EncryptedData"`
               } `xml:"LA"`
               Signature struct {
                  SignedInfo struct {
                     Xmlns                  string `xml:"xmlns,attr"`
                     CanonicalizationMethod struct {
                        Algorithm string `xml:"Algorithm,attr"`
                     } `xml:"CanonicalizationMethod"`
                     SignatureMethod struct {
                        Algorithm string `xml:"Algorithm,attr"`
                     } `xml:"SignatureMethod"`
                     Reference struct {
                        Uri          string `xml:"URI,attr"`
                        DigestMethod struct {
                           Algorithm string `xml:"Algorithm,attr"`
                        } `xml:"DigestMethod"`
                        DigestValue string `xml:"DigestValue"`
                     } `xml:"Reference"`
                  } `xml:"SignedInfo"`
                  SignatureValue string `xml:"SignatureValue"`
               } `xml:"Signature"`
            } `xml:"Challenge"`
         } `xml:"challenge"`
      } `xml:"AcquireLicense"`
   } `xml:"soap:Body"`
}
