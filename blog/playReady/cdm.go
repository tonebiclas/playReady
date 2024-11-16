package playReady

import (
)

type cdm struct{}

func (cdm) license_challenge(session_id []byte, content_header string) string {
  main_body = (
      '<?xml version="1.0" encoding="utf-8"?>'
      '<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">'
          '<soap:Body>'
              '<AcquireLicense xmlns="http://schemas.microsoft.com/DRM/2007/03/protocols">'
                  '<challenge>'
                      '<Challenge xmlns="http://schemas.microsoft.com/DRM/2007/03/protocols/messages">'
                          + la_content +
                          '<Signature xmlns="http://www.w3.org/2000/09/xmldsig#">'
                              + signed_info +
                              f'<SignatureValue>{base64.b64encode(signature).decode()}</SignatureValue>'
                              '<KeyInfo xmlns="http://www.w3.org/2000/09/xmldsig#">'
                                  '<KeyValue>'
                                      '<ECCKeyValue>'
                                          f'<PublicKey>{base64.b64encode(session.signing_key.public_bytes()).decode()}</PublicKey>'
                                      '</ECCKeyValue>'
                                  '</KeyValue>'
                              '</KeyInfo>'
                          '</Signature>'
                      '</Challenge>'
                  '</challenge>'
              '</AcquireLicense>'
          '</soap:Body>'
      '</soap:Envelope>'
  )
  return main_body
}

type Envelope struct {
   XMLName xml.Name `xml:"Envelope"`
   Text    string   `xml:",chardata"`
   Soap    string   `xml:"soap,attr"`
   Body    struct {
      Text           string `xml:",chardata"`
      AcquireLicense struct {
         Text      string `xml:",chardata"`
         Xmlns     string `xml:"xmlns,attr"`
         Challenge struct {
            Text      string `xml:",chardata"`
            Challenge struct {
               Text  string `xml:",chardata"`
               Xmlns string `xml:"xmlns,attr"`
               LA    struct {
                  Text          string `xml:",chardata"`
                  Xmlns         string `xml:"xmlns,attr"`
                  ID            string `xml:"Id,attr"`
                  Space         string `xml:"space,attr"`
                  Version       string `xml:"Version"`
                  ContentHeader struct {
                     Text      string `xml:",chardata"`
                     WRMHEADER struct {
                        Text    string `xml:",chardata"`
                        Xmlns   string `xml:"xmlns,attr"`
                        Version string `xml:"version,attr"`
                        DATA    struct {
                           Text        string `xml:",chardata"`
                           PROTECTINFO struct {
                              Text   string `xml:",chardata"`
                              KEYLEN string `xml:"KEYLEN"`
                              ALGID  string `xml:"ALGID"`
                           } `xml:"PROTECTINFO"`
                           KID string `xml:"KID"`
                        } `xml:"DATA"`
                     } `xml:"WRMHEADER"`
                  } `xml:"ContentHeader"`
                  CLIENTINFO struct {
                     Text          string `xml:",chardata"`
                     CLIENTVERSION string `xml:"CLIENTVERSION"`
                  } `xml:"CLIENTINFO"`
                  RevocationLists struct {
                     Text        string `xml:",chardata"`
                     RevListInfo []struct {
                        Text    string `xml:",chardata"`
                        ListID  string `xml:"ListID"`
                        Version string `xml:"Version"`
                     } `xml:"RevListInfo"`
                  } `xml:"RevocationLists"`
                  LicenseNonce  string `xml:"LicenseNonce"`
                  ClientTime    string `xml:"ClientTime"`
                  EncryptedData struct {
                     Text             string `xml:",chardata"`
                     Xmlns            string `xml:"xmlns,attr"`
                     Type             string `xml:"Type,attr"`
                     EncryptionMethod struct {
                        Text      string `xml:",chardata"`
                        Algorithm string `xml:"Algorithm,attr"`
                     } `xml:"EncryptionMethod"`
                     KeyInfo struct {
                        Text         string `xml:",chardata"`
                        Xmlns        string `xml:"xmlns,attr"`
                        EncryptedKey struct {
                           Text             string `xml:",chardata"`
                           Xmlns            string `xml:"xmlns,attr"`
                           EncryptionMethod struct {
                              Text      string `xml:",chardata"`
                              Algorithm string `xml:"Algorithm,attr"`
                           } `xml:"EncryptionMethod"`
                           KeyInfo struct {
                              Text    string `xml:",chardata"`
                              Xmlns   string `xml:"xmlns,attr"`
                              KeyName string `xml:"KeyName"`
                           } `xml:"KeyInfo"`
                           CipherData struct {
                              Text        string `xml:",chardata"`
                              CipherValue string `xml:"CipherValue"`
                           } `xml:"CipherData"`
                        } `xml:"EncryptedKey"`
                     } `xml:"KeyInfo"`
                     CipherData struct {
                        Text        string `xml:",chardata"`
                        CipherValue string `xml:"CipherValue"`
                     } `xml:"CipherData"`
                  } `xml:"EncryptedData"`
               } `xml:"LA"`
               Signature struct {
                  Text       string `xml:",chardata"`
                  SignedInfo struct {
                     Text                   string `xml:",chardata"`
                     Xmlns                  string `xml:"xmlns,attr"`
                     CanonicalizationMethod struct {
                        Text      string `xml:",chardata"`
                        Algorithm string `xml:"Algorithm,attr"`
                     } `xml:"CanonicalizationMethod"`
                     SignatureMethod struct {
                        Text      string `xml:",chardata"`
                        Algorithm string `xml:"Algorithm,attr"`
                     } `xml:"SignatureMethod"`
                     Reference struct {
                        Text         string `xml:",chardata"`
                        URI          string `xml:"URI,attr"`
                        DigestMethod struct {
                           Text      string `xml:",chardata"`
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
   } `xml:"Body"`
} 
