package config

//Constants
const (
	CertSeedWdsl = "https://maullin.sii.cl/DTEWS/CrSeed.jws?WSDL"
	ProdSeedWdsl = "https://palena.sii.cl/DTEWS/CrSeed.jws?WSDL"
	SeedTemplate = `
	<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:def="http://DefaultNamespace">
		<soapenv:Header/>
		<soapenv:Body>
			<def:getSeed soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"/>
		</soapenv:Body>
	</soapenv:Envelope>`
	CertTokenWsdl = "https://maullin.sii.cl/DTEWS/GetTokenFromSeed.jws?WSDL"
	ProdTokenWsdl = "https://maullin.sii.cl/DTEWS/GetTokenFromSeed.jws?WSDL"
	TokenTemplate = `
	<soapenv:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:def="http://DefaultNamespace">
		<soapenv:Header/>
		<soapenv:Body>
			<def:getToken soapenv:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/">
				<pszXml xsi:type="xsd:string"><![CDATA[@pszXML]]></pszXml>
			</def:getToken>
		</soapenv:Body>
	</soapenv:Envelope>`
	PszXML = `
	<getToken>
		<item>
			<Semilla>@seed</Semilla>
		</item>
		<Signature xmlns="http://www.w3.org/2000/09/xmldsig#">
			<SignedInfo>
				<CanonicalizationMethod Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315"/>
				<SignatureMethod Algorithm="http://www.w3.org/2000/09/xmldsig#rsa-sha1"/>
				<Reference URI="">
					<Transforms>
						<Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature"/>
					</Transforms>
					<DigestMethod Algorithm="http://www.w3.org/2000/09/xmldsig#sha1"/>
					<DigestValue></DigestValue>
				</Reference>
			</SignedInfo>
			<SignatureValue/>
			<KeyInfo>
				<KeyValue/>
				<X509Data><X509Certificate/></X509Data>
			</KeyInfo>
		</Signature>
	</getToken>`
	PurchaseDetailURL = "https://www4.sii.cl/consdcvinternetui/services/data/facadeService/getDetalleCompraExport"
)
