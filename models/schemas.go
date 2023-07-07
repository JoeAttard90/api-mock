
package models


type ClientFundDataResponse struct {
    ExistingDocumentCount int `json:"existingDocumentCount"`
    FileName string `json:"fileName"`
    IssuerId string `json:"issuerId"`
    AccessBusinessLineName []string `json:"accessBusinessLineName"`
    Exists bool `json:"exists"`
    IssuerName string `json:"issuerName"`
    ExistingDocumentVersion []string `json:"existingDocumentVersion"`
    AssetID string `json:"assetID"`
    AssetName string `json:"assetName"`
    BusinessLineMEName []string `json:"businessLineMEName"`
    Errors []string `json:"errors"`
    ExistingDocumentStatus []string `json:"existingDocumentStatus"`
    BusinessLineMEID []string `json:"businessLineMEID"`
    RequestId string `json:"requestId"`
}

type ClientUserDetailsResponse struct {
    BusinessLines any `json:"businessLines"`
    ClientUserPhone string `json:"clientUserPhone"`
    ClientUserNotificationEmail2 string `json:"clientUserNotificationEmail2"`
    Errors []string `json:"errors"`
    RequestId string `json:"requestId"`
    ClientUserTitle string `json:"clientUserTitle"`
    ManagedEntity any `json:"managedEntity"`
    ClientUserCountry string `json:"clientUserCountry"`
    ClientUserEmail string `json:"clientUserEmail"`
    ClientUserLastName string `json:"clientUserLastName"`
    ClientUserRole []string `json:"clientUserRole"`
    ClientUserFirstName string `json:"clientUserFirstName"`
    ClientUserMiddleInitials string `json:"clientUserMiddleInitials"`
    ClientUserNotificationEmail1 string `json:"clientUserNotificationEmail1"`
    RelationshipData any `json:"relationshipData"`
}

type RelationshipData struct {
    BusinessUserFirstName string `json:"businessUserFirstName"`
    BusinessUserJobTitle string `json:"businessUserJobTitle"`
    BusinessUserLastName string `json:"businessUserLastName"`
    BusinessUserPortalEmail string `json:"businessUserPortalEmail"`
    BusinessUserRole string `json:"businessUserRole"`
    BusinessUserTeam string `json:"businessUserTeam"`
    MeRelationships any `json:"meRelationships"`
}

type SearchClientUserDocsRequest struct {
    Page int `json:"page"`
    PageSize int `json:"pageSize"`
    SearchToApply any `json:"searchToApply"`
    SortField any `json:"sortField"`
    FiltersToApply any `json:"filtersToApply"`
}

type Assets struct {
    AssetID int `json:"assetID"`
    AssetName string `json:"assetName"`
    IssuerName string `json:"issuerName"`
    SourceIssuerID int `json:"sourceIssuerID"`
}

type BusinessUserDetailsResponse struct {
    AccessData any `json:"accessData"`
    BusinessUserEmail string `json:"businessUserEmail"`
    BusinessUserID int `json:"businessUserID"`
    BusinessUserJobTitle string `json:"businessUserJobTitle"`
    BusinessUserLastName string `json:"businessUserLastName"`
    BusinessUserTeam string `json:"businessUserTeam"`
    BusinessUserFirstName string `json:"businessUserFirstName"`
    BusinessUserMiddleInitials string `json:"businessUserMiddleInitials"`
    BusinessUserTitle string `json:"businessUserTitle"`
    Errors []string `json:"errors"`
    FilterSummaries any `json:"filterSummaries"`
    RequestId string `json:"requestId"`
}

type MeRelationship struct {
    MeRelationshipId int `json:"meRelationshipId"`
    MeRelationshipName string `json:"meRelationshipName"`
}

type PostMetaDataResponseBusinessUserRejection struct {
    BusinessLineIdTag string `json:"businessLineIdTag"`
    Errors []string `json:"errors"`
    RejectionEmail any `json:"rejectionEmail"`
    RequestId string `json:"requestId"`
}

type PostMetaDataResponseClientUser struct {
    BusinessLineIdTag string `json:"businessLineIdTag"`
    ClientUsers any `json:"clientUsers"`
    Errors []string `json:"errors"`
    RequestId string `json:"requestId"`
}

type Term struct {
    Count int `json:"count"`
    Label string `json:"label"`
    Value string `json:"value"`
}

type ClientUserDocsEntitlement struct {
    AssetClass string `json:"assetClass"`
    BusinessLineId int `json:"businessLineId"`
    BusinessLineShortName string `json:"businessLineShortName"`
    DocumentLinkage string `json:"documentLinkage"`
    DocumentMonth string `json:"documentMonth"`
    DocumentReadStatus string `json:"documentReadStatus"`
    DocumentUpdatedDate string `json:"documentUpdatedDate"`
    IsConfidential string `json:"isConfidential"`
    ManagedEntityId []int `json:"managedEntityId"`
    AssetId []int `json:"assetId"`
    Audience string `json:"audience"`
    DocumentEffectiveDate string `json:"documentEffectiveDate"`
    DocumentFileName string `json:"documentFileName"`
    DocumentSubCategory string `json:"documentSubCategory"`
    DocumentType string `json:"documentType"`
    DocumentYear int `json:"documentYear"`
    BusinessLineName string `json:"businessLineName"`
    DocumentCategory string `json:"documentCategory"`
    DocumentID int `json:"documentID"`
    FileSize int `json:"fileSize"`
    Capability string `json:"capability"`
    DocumentQuarter string `json:"documentQuarter"`
    DocumentVersion string `json:"documentVersion"`
    Format string `json:"format"`
}

type ClientUserDocsResponse struct {
    RequestId string `json:"requestId"`
    TotalPages int `json:"totalPages"`
    TotalResults int `json:"totalResults"`
    Assets any `json:"assets"`
    Documents any `json:"documents"`
    Errors []string `json:"errors"`
    FilterSummaries any `json:"filterSummaries"`
    ManagedEntities any `json:"managedEntities"`
}

type Search struct {
    Value string `json:"value"`
}

type BusinessUserDocsEntitlement struct {
    AssetId []int `json:"assetId"`
    Audience string `json:"audience"`
    DocumentMonth string `json:"documentMonth"`
    DocumentStatus string `json:"documentStatus"`
    RejectReason string `json:"rejectReason"`
    UploadedByUserId int `json:"uploadedByUserId"`
    AssetName string `json:"assetName"`
    BusinessLineName string `json:"businessLineName"`
    DocumentCategory string `json:"documentCategory"`
    UploadedByUser string `json:"uploadedByUser"`
    DocumentLastUpdatedDate string `json:"DocumentLastUpdatedDate"`
    Capability string `json:"capability"`
    DocumentLinkage string `json:"documentLinkage"`
    IsConfidential string `json:"isConfidential"`
    ManagedEntityId []int `json:"managedEntityId"`
    BusinessLineShortName string `json:"businessLineShortName"`
    DocumentYear int `json:"documentYear"`
    ExpiryDate string `json:"expiryDate"`
    ApprovedDate string `json:"approvedDate"`
    DocumentQuarter string `json:"documentQuarter"`
    Format string `json:"format"`
    UploadedByUserEmail string `json:"uploadedByUserEmail"`
    ApprovedByUser string `json:"approvedByUser"`
    AssetClass string `json:"assetClass"`
    RejectedDate string `json:"rejectedDate"`
    DocumentSubCategory string `json:"documentSubCategory"`
    DocumentType string `json:"documentType"`
    UploadedDate string `json:"uploadedDate"`
    DocumentVersion string `json:"documentVersion"`
    FileSize int `json:"fileSize"`
    RejectedByUser string `json:"rejectedByUser"`
    DocumentEffectiveDate string `json:"documentEffectiveDate"`
    DocumentFileName string `json:"documentFileName"`
    DocumentID int `json:"documentID"`
}

type PostMetaDataBusinessUser struct {
    Email string `json:"email"`
}

type PostMetaDataClientUser struct {
    ClientUserFirstName string `json:"clientUserFirstName"`
    ClientUserLastName string `json:"clientUserLastName"`
    ClientUserMiddleInitials string `json:"clientUserMiddleInitials"`
    ClientUserNotificationEmail1 string `json:"clientUserNotificationEmail1"`
    ClientUserNotificationEmail2 string `json:"clientUserNotificationEmail2"`
    ClientUserTitle string `json:"clientUserTitle"`
}

type ManagedEntity struct {
    InvestorGroupName string `json:"InvestorGroupName"`
    ManagedEntityId int `json:"managedEntityId"`
    ManagedEntityName string `json:"managedEntityName"`
}

type ManagedEntityResponse struct {
    ManagedEntityId int `json:"managedEntityId"`
    ManagedEntityName string `json:"managedEntityName"`
    ManagedEntityBusinessLines any `json:"managedEntityBusinessLines"`
}

type UploadEntitlementResponse struct {
    Errors []string `json:"errors"`
    IsEntitled bool `json:"isEntitled"`
    RequestId string `json:"requestId"`
}

type DocEntitlementResponse struct {
    IsEntitled bool `json:"isEntitled"`
    RequestId string `json:"requestId"`
    Version string `json:"version"`
    Errors []string `json:"errors"`
    FileName string `json:"fileName"`
}

type SortField struct {
    Direction string `json:"direction"`
    Field string `json:"field"`
}

type Filter struct {
    Field string `json:"field"`
    Value string `json:"value"`
}

type FilterSummary struct {
    Field string `json:"field"`
    Terms any `json:"terms"`
}

type ManagedEntityBusinessLineResponse struct {
    BusinessLineMEName string `json:"businessLineMEName"`
    ManagedEntityBusinessLineShortName string `json:"managedEntityBusinessLineShortName"`
    ValuationId string `json:"valuationId"`
    BusinessLineMEID string `json:"businessLineMEID"`
}

type AccessData struct {
    AccessBusinessLineShortName string `json:"accessBusinessLineShortName"`
    AccessRole string `json:"accessRole"`
    AdGroup string `json:"adGroup"`
    AccessBusinessLine string `json:"accessBusinessLine"`
    AccessBusinessLineID int `json:"accessBusinessLineID"`
}

type BusinessUserDocsResponse struct {
    FilterSummaries any `json:"filterSummaries"`
    ManagedEntities any `json:"managedEntities"`
    RequestId string `json:"requestId"`
    TotalPages int `json:"totalPages"`
    TotalResults int `json:"totalResults"`
    Assets any `json:"assets"`
    Documents any `json:"documents"`
    Errors []string `json:"errors"`
}

type PostMetaDataResponseBusinessUserApproval struct {
    ApprovalEmail any `json:"approvalEmail"`
    BusinessLineIdTag string `json:"businessLineIdTag"`
    Errors []string `json:"errors"`
    RequestId string `json:"requestId"`
}

type PostReadStatusRequest struct {
    FileName string `json:"fileName"`
}

type Response struct {
    Errors []string `json:"errors"`
    RequestId string `json:"requestId"`
}

type SearchBusinessUserDocsRequest struct {
    Page int `json:"page"`
    PageSize int `json:"pageSize"`
    SearchToApply any `json:"searchToApply"`
    SortField any `json:"sortField"`
    FiltersToApply any `json:"filtersToApply"`
}

type BusinessLine struct {
    BusinessLineId int `json:"businessLineId"`
    BusinessLineName string `json:"businessLineName"`
    Capability string `json:"Capability"`
    AssetClass string `json:"assetClass"`
}

type PostMetaDataRequest struct {
    BusinessLineMEID []string `json:"businessLineMEID"`
    BusinessLineShortName string `json:"businessLineShortName"`
    DocumentMonth string `json:"documentMonth"`
    DocumentYear int `json:"documentYear"`
    PreviousDocStatus string `json:"previousDocStatus"`
    Version string `json:"version"`
    Action string `json:"action"`
    AssetID string `json:"assetID"`
    DocumentCategory string `json:"documentCategory"`
    DocumentLinkage string `json:"documentLinkage"`
    DocumentStatus string `json:"documentStatus"`
    DocumentSubCategory string `json:"documentSubCategory"`
    FileSize int `json:"fileSize"`
    IsConfidential string `json:"isConfidential"`
    UpdatedByUser string `json:"updatedByUser"`
    ExpiryDate string `json:"expiryDate"`
    FileName string `json:"fileName"`
    Format string `json:"format"`
    RejectReason string `json:"rejectReason"`
    UpdatedDate string `json:"updatedDate"`
    AccessBusinessLineName []string `json:"accessBusinessLineName"`
    Audience string `json:"audience"`
    DocumentQuarter string `json:"documentQuarter"`
    DocumentType string `json:"documentType"`
    EffectiveDate string `json:"effectiveDate"`
    IssuerId string `json:"issuerId"`
}

