package template_variables

// Please, DO NOT EDIT CONST
// just add
const (
	/*

			**********************************************************************
		                                    STATUSES
			**********************************************************************

	*/

	// // -- type 5
	NewEntityStatusTypeFive                    = "6281eebf21985beac89b5a85" //new status_id for type_code=5 entity
	InAuctionTypeFive                          = "62f4eecfe0506943910a48db" // Auksion sotuvida
	SoldTypeFive                               = "62a6cc37a92f4e923db6cd41"
	DavreestrFixIssuesTypeFive                 = "6283c45b0effe0b32f2ac3da" // Davreestr xatoliklarini tahrirlash [Y]
	InDiscussionTypeFive                       = ""                         //
	InDiscussionProcessFive                    = "62610bb0b06e4c2cc493086e"
	RejectedTypeFive                           = "62610d89b06e4c2cc4930a18"
	AuctionReceivedTypeFive                    = "63187c21bdd563bac21131fc" //
	SentToAuctionTypeFive                      = "62624ad787cae4b89c431d9f" // Auksionga yuborilgan
	DavreestrWinnerRegistrationErrorStatusFive = "62e928e8ade0ea57cf20117e" // G'olib nomiga ro'yxatdan o'tkazishda xatoliklarni tahrirlash
	DecisionStage                              = "6261042cb06e4c2cc492febf"
	RegisterNewEntityTypeFive                  = "6261049db06e4c2cc492ff53"
	GotRegisteredStageTypeFive                 = "62a6cd13a92f4e923db6ce56"
	SendToKelishuvTypeFive                     = "627baa176a9eaec8edca8f84"
	FixMistakesTypeFive                        = "631f41f208ebb7c0210171ef" // Xatoliklarni tahrirlash

	// // -- type 6
	DSQevaluatePriceTypeSix                = "61fa619ca9a0964dbfcdd599" // DSQ da boshlang'ich narxni belgilash
	NewEntityStatusTypeSix                 = "61fa5d7ea9a0964dbfcdd58d" //new status_id for type_code=6 entity
	RegionProtectionZoneCheck              = "625660fbd126bdf4e54eb056" // Viloyatda himoya hududini tekshirish
	FixMistakesTypeSix                     = "62331cf8e936be852bfec01b"
	RejectedTypeSix                        = "61f24c07c105e1d57c2bd047"
	AuctionReceivedTypeSix                 = "627e8ee75d7af0cc3032f3d9"
	SentToAuctionTypeSix                   = "61fa61c8a9a0964dbfcdd59a"
	RejectLandFromWinnerTypeSix            = "63353f17d23d0b27354b572e"
	Rejected2TypeSix                       = "62b18a3fe50f7a3d44f622df"
	ResendToAuctionRypeSix                 = "62331d6be936be852bfec02c"
	InAuctionTypeSix                       = "61fa621da9a0964dbfcdd59c"
	SoldTypeSix                            = "61fa624fa9a0964dbfcdd59d"
	InDiscussionTypeSix                    = "61fa60b2a9a0964dbfcdd594"
	InDiscussionProcessSix                 = "61fa6019a9a0964dbfcdd592"
	DavreestrFixIssuesTypeSix              = "6283c3fe0effe0b32f2ac3aa" // Davreestr xatoliklarini tahrirlash
	DavreestrWinnerRegistrationErrorStatus = "62be78e1542bd77ff4b0b8b7" // G'olib nomiga ro'yxatdan o'tkazishda xatoliklarni tahrirlash
	EntityMaterialChoosingStage            = "61fa5e8aa9a0964dbfcdd58f"
	EntityLoyihalashdaStatusIdTypeSix      = "61fa5f6fa9a0964dbfcdd590"
	DocumentGetheringStage                 = "61f268a2c105e1d57c2bd072"
	SendToRavreestrStageTypeSix            = "6299ae061d86a5754e68bd74"
	RegisterNewEntityStageSix              = "62038f2b29a37eaa494fcf63"
	IdentifyPolicyTypeStage                = "61fa616fa9a0964dbfcdd598"
	GotRegisteredStage                     = "61fa629ba9a0964dbfcdd59e"
	SendToKelishuvTypeSix                  = "61fa5fe5a9a0964dbfcdd591"
	SendToMuhokamaTypeSix                  = "61fa6049a9a0964dbfcdd593"
	DavaktivStartingPriceTypeSix           = "61fb6a26d007230eab60d6e8"
	MakeOfferTypeSix                       = "61fa5dc8a9a0964dbfcdd58e"

	// // -- type draft
	NewEntityDraftStatus = "6103887b74795f82ba173572" //new status_id for entityDraft
	RejectDraftStatusID  = "620b886f9b6106043c42f4a5"

	/*

			**********************************************************************
		                                    PROPERTIES
			**********************************************************************

	*/

	/*

			**********************************************************************
		                                    ORGANIZATIONS
			**********************************************************************

	*/
	UzGASHLITI = "61fa7a2b9cc8808579eb9930"

		/*

			**********************************************************************
		                                    URLS AND AUTH
			**********************************************************************

	*/
	// DavreestrUrls
	DavReestrUrl          = "https://davreestr.kadastr.uz/api/yer-elektron/push"

	// Soliq
	SoliqLoginURL = "http://10.190.27.58:8103/reestr-user-api/anonymous/login"
	SoliqURL      = "http://10.190.27.58:8103/reestr-admin-api/company/get/"
	SoliqUsername = "XXXXX"
	SoliqPassword = "YYYYY"
)

var UpdateForbiddenStatuses = []string{
	InDiscussionProcessSix,
	InDiscussionTypeSix,
	FixMistakesTypeSix,
	DavreestrFixIssuesTypeSix,
}

var SaveButtonForbiddenStatuses = []string{
	SendToKelishuvTypeSix,
	InDiscussionProcessSix,
	SendToMuhokamaTypeSix,
	InDiscussionTypeSix,
}

var ProjectOrganizationDefaultPermissions = []string{
	"61028a003ed0c289c2eecbd3",
	"61028a003ed0c289c2eecbd5",
	"61028a003ed0c289c2eecbd6",
	"6135dc19a55f7420e58bf806",
	"6142ed926ff4d84171ad75f5",
}

var ConstructionTypesWithAuctionCode = map[string]string{
	"6142c2046074f7fa21292a28": "1",
	"6142c2046074f7fa21292a29": "2",
	"6142c2056074f7fa21292a2a": "2.1",
	"6142c2056074f7fa21292a2b": "3",
	"6142c2056074f7fa21292a2c": "3.1",
	"6142c2056074f7fa21292a2d": "4",
	"6142c2056074f7fa21292a2e": "4.1",
	"6142c2056074f7fa21292a2f": "4.2",
	"6142c2066074f7fa21292a30": "4.3",
	"6142c2066074f7fa21292a31": "4.4",
	"6142c2066074f7fa21292a32": "5",
	"6142c2066074f7fa21292a33": "5.1",
	"6142c2066074f7fa21292a34": "6",
	"6142c2066074f7fa21292a35": "6.1",
	"6142c2066074f7fa21292a36": "7",
	"6142c2076074f7fa21292a37": "7.1",
	"6142c2076074f7fa21292a38": "7.1.1",
	"6142c2076074f7fa21292a39": "8",
	"6142c2076074f7fa21292a3a": "8.1",
	"6142c2076074f7fa21292a3b": "9",
	"6142c2076074f7fa21292a3c": "9.1",
	"6142c2086074f7fa21292a3d": "9.1.1",
	"6142c2086074f7fa21292a3e": "9.1.2",
	"6142c2146074f7fa21292a87": "9.1.3",
	"6242e5ce748e61558f1de6b5": "9.1.4",
	"624d8082ed61d1fa13ae2221": "9.1.5",
	"6142c2086074f7fa21292a3f": "10",
	"6142c2086074f7fa21292a40": "10.1",
	"6142c2086074f7fa21292a41": "10.1.1",
	"6142c2086074f7fa21292a42": "10.2",
	"6142c2086074f7fa21292a43": "11",
	"6142c2096074f7fa21292a44": "11.1",
	"6142c2096074f7fa21292a45": "12",
	"6142c2096074f7fa21292a46": "12.1",
	"6142c2096074f7fa21292a47": "12.2",
	"6142c2096074f7fa21292a48": "12.3",
	"6142c2096074f7fa21292a49": "13",
	"6142c20a6074f7fa21292a4a": "13.1",
	"6142c20a6074f7fa21292a4b": "13.2",
	"6142c20a6074f7fa21292a4c": "13.3",
	"6142c20a6074f7fa21292a4d": "13.4",
	"6142c20a6074f7fa21292a4e": "13.5",
	"6142c20a6074f7fa21292a4f": "14",
	"6142c20b6074f7fa21292a50": "14.1",
	"6142c20b6074f7fa21292a51": "15",
	"6142c20b6074f7fa21292a52": "15.1",
	"6142c20b6074f7fa21292a53": "15.1.1",
	"6142c20b6074f7fa21292a54": "15.1.2",
	"6142c20c6074f7fa21292a55": "15.2",
	"6142c20c6074f7fa21292a56": "15.2.1",
	"6142c20c6074f7fa21292a57": "15.2.2",
	"6142c20c6074f7fa21292a58": "16",
	"6142c20c6074f7fa21292a59": "16.1",
	"6142c20c6074f7fa21292a5a": "16.2",
	"6142c20c6074f7fa21292a5b": "17",
	"6142c20d6074f7fa21292a5c": "17.1",
	"6142c20d6074f7fa21292a5d": "17.1.1",
	"6142c20d6074f7fa21292a5e": "17.1.2",
	"6142c20d6074f7fa21292a5f": "17.1.3",
	"6142c20d6074f7fa21292a60": "17.2",
	"6142c20e6074f7fa21292a61": "17.3",
	"6142c20e6074f7fa21292a62": "18",
	"6142c20e6074f7fa21292a63": "18.1",
	"6142c20e6074f7fa21292a64": "18.2",
	"6142c20e6074f7fa21292a65": "18.2.1",
	"6142c20e6074f7fa21292a66": "18.2.2",
	"6142c20f6074f7fa21292a67": "18.3",
	"6142c20f6074f7fa21292a68": "19",
	"6142c20f6074f7fa21292a69": "19.1",
	"6142c20f6074f7fa21292a6a": "19.1.1",
	"6142c20f6074f7fa21292a6b": "20",
	"6142c20f6074f7fa21292a6c": "20.1",
	"6142c2106074f7fa21292a6d": "20.2",
	"6142c2106074f7fa21292a6e": "20.3",
	"6142c2106074f7fa21292a6f": "20.4",
	"6142c2106074f7fa21292a70": "20.5",
	"639c3af01ddaf2be4813d945": "20.6",
	"6142c2106074f7fa21292a71": "21",
	"6142c2106074f7fa21292a72": "21.1",
	"6142c2116074f7fa21292a73": "21.2",
	"6142c2116074f7fa21292a74": "21.3",
	"6142c2116074f7fa21292a75": "21.4",
	"6142c2116074f7fa21292a76": "22",
	"6142c2116074f7fa21292a77": "22.1",
	"6142c2116074f7fa21292a78": "22.2",
	"6142c2126074f7fa21292a79": "22.3",
	"6142c2126074f7fa21292a7a": "23",
	"6142c2126074f7fa21292a7b": "23.1",
	"6142c2126074f7fa21292a7c": "23.1.1",
	"6142c2126074f7fa21292a7d": "23.1.2",
	"6142c2126074f7fa21292a7e": "24",
	"6142c2136074f7fa21292a7f": "24.1",
	"6142c2136074f7fa21292a80": "24.2",
	"6142c2136074f7fa21292a81": "25",
	"6142c2136074f7fa21292a82": "25.1",
	"6142c2146074f7fa21292a83": "25.2",
	"6142c2146074f7fa21292a84": "25.3",
	"6142c2146074f7fa21292a85": "25.4",
	"6142c2146074f7fa21292a86": "25.5",
	"632bfec15732f0984f1cddeb": "26",
}
