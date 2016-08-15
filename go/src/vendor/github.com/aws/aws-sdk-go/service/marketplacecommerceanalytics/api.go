// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package marketplacecommerceanalytics provides a client for AWS Marketplace Commerce Analytics.
package marketplacecommerceanalytics

import (
	"time"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opGenerateDataSet = "GenerateDataSet"

// GenerateDataSetRequest generates a "aws/request.Request" representing the
// client's request for the GenerateDataSet operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the GenerateDataSet method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the GenerateDataSetRequest method.
//    req, resp := client.GenerateDataSetRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
func (c *MarketplaceCommerceAnalytics) GenerateDataSetRequest(input *GenerateDataSetInput) (req *request.Request, output *GenerateDataSetOutput) {
	op := &request.Operation{
		Name:       opGenerateDataSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GenerateDataSetInput{}
	}

	req = c.newRequest(op, input, output)
	output = &GenerateDataSetOutput{}
	req.Data = output
	return
}

// Given a data set type and data set publication date, asynchronously publishes
// the requested data set to the specified S3 bucket and notifies the specified
// SNS topic once the data is available. Returns a unique request identifier
// that can be used to correlate requests with notifications from the SNS topic.
// Data sets will be published in comma-separated values (CSV) format with the
// file name {data_set_type}_YYYY-MM-DD.csv. If a file with the same name already
// exists (e.g. if the same data set is requested twice), the original file
// will be overwritten by the new file. Requires a Role with an attached permissions
// policy providing Allow permissions for the following actions: s3:PutObject,
// s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish, iam:GetRolePolicy.
func (c *MarketplaceCommerceAnalytics) GenerateDataSet(input *GenerateDataSetInput) (*GenerateDataSetOutput, error) {
	req, out := c.GenerateDataSetRequest(input)
	err := req.Send()
	return out, err
}

const opStartSupportDataExport = "StartSupportDataExport"

// StartSupportDataExportRequest generates a "aws/request.Request" representing the
// client's request for the StartSupportDataExport operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the StartSupportDataExport method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the StartSupportDataExportRequest method.
//    req, resp := client.StartSupportDataExportRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
func (c *MarketplaceCommerceAnalytics) StartSupportDataExportRequest(input *StartSupportDataExportInput) (req *request.Request, output *StartSupportDataExportOutput) {
	op := &request.Operation{
		Name:       opStartSupportDataExport,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartSupportDataExportInput{}
	}

	req = c.newRequest(op, input, output)
	output = &StartSupportDataExportOutput{}
	req.Data = output
	return
}

// Given a data set type and a from date, asynchronously publishes the requested
// customer support data to the specified S3 bucket and notifies the specified
// SNS topic once the data is available. Returns a unique request identifier
// that can be used to correlate requests with notifications from the SNS topic.
// Data sets will be published in comma-separated values (CSV) format with the
// file name {data_set_type}_YYYY-MM-DD'T'HH-mm-ss'Z'.csv. If a file with the
// same name already exists (e.g. if the same data set is requested twice),
// the original file will be overwritten by the new file. Requires a Role with
// an attached permissions policy providing Allow permissions for the following
// actions: s3:PutObject, s3:GetBucketLocation, sns:GetTopicAttributes, sns:Publish,
// iam:GetRolePolicy.
func (c *MarketplaceCommerceAnalytics) StartSupportDataExport(input *StartSupportDataExportInput) (*StartSupportDataExportOutput, error) {
	req, out := c.StartSupportDataExportRequest(input)
	err := req.Send()
	return out, err
}

// Container for the parameters to the GenerateDataSet operation.
type GenerateDataSetInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Key-value pairs which will be returned, unmodified, in the Amazon
	// SNS notification message and the data set metadata file. These key-value
	// pairs can be used to correlated responses with tracking information from
	// other systems.
	CustomerDefinedValues map[string]*string `locationName:"customerDefinedValues" min:"1" type:"map"`

	// The date a data set was published. For daily data sets, provide a date with
	// day-level granularity for the desired day. For weekly data sets, provide
	// a date with day-level granularity within the desired week (the day value
	// will be ignored). For monthly data sets, provide a date with month-level
	// granularity for the desired month (the day value will be ignored).
	DataSetPublicationDate *time.Time `locationName:"dataSetPublicationDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The desired data set type.
	//
	//   customer_subscriber_hourly_monthly_subscriptions - Available daily by
	// 5:00 PM Pacific Time since 2014-07-21. customer_subscriber_annual_subscriptions
	// - Available daily by 5:00 PM Pacific Time since 2014-07-21. daily_business_usage_by_instance_type
	// - Available daily by 5:00 PM Pacific Time since 2015-01-26. daily_business_fees
	// - Available daily by 5:00 PM Pacific Time since 2015-01-26. daily_business_free_trial_conversions
	// - Available daily by 5:00 PM Pacific Time since 2015-01-26. daily_business_new_instances
	// - Available daily by 5:00 PM Pacific Time since 2015-01-26. daily_business_new_product_subscribers
	// - Available daily by 5:00 PM Pacific Time since 2015-01-26. daily_business_canceled_product_subscribers
	// - Available daily by 5:00 PM Pacific Time since 2015-01-26. monthly_revenue_billing_and_revenue_data
	// - Available monthly on the 4th day of the month by 5:00 PM Pacific Time since
	// 2015-02. monthly_revenue_annual_subscriptions - Available monthly on the
	// 4th day of the month by 5:00 PM Pacific Time since 2015-02. disbursed_amount_by_product
	// - Available every 30 days by 5:00 PM Pacific Time since 2015-01-26. disbursed_amount_by_product_with_uncollected_funds
	// -This data set is only available from 2012-04-19 until 2015-01-25. After
	// 2015-01-25, this data set was split into three data sets: disbursed_amount_by_product,
	// disbursed_amount_by_age_of_uncollected_funds, and disbursed_amount_by_age_of_disbursed_funds.
	//  disbursed_amount_by_customer_geo - Available every 30 days by 5:00 PM Pacific
	// Time since 2012-04-19. disbursed_amount_by_age_of_uncollected_funds - Available
	// every 30 days by 5:00 PM Pacific Time since 2015-01-26. disbursed_amount_by_age_of_disbursed_funds
	// - Available every 30 days by 5:00 PM Pacific Time since 2015-01-26. customer_profile_by_industry
	// - Available daily by 5:00 PM Pacific Time since 2015-10-01. customer_profile_by_revenue
	// - Available daily by 5:00 PM Pacific Time since 2015-10-01. customer_profile_by_geography
	// - Available daily by 5:00 PM Pacific Time since 2015-10-01.
	DataSetType *string `locationName:"dataSetType" min:"1" type:"string" required:"true" enum:"DataSetType"`

	// The name (friendly name, not ARN) of the destination S3 bucket.
	DestinationS3BucketName *string `locationName:"destinationS3BucketName" min:"1" type:"string" required:"true"`

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket
	// name "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If
	// the prefix directory structure does not exist, it will be created. If no
	// prefix is provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string `locationName:"destinationS3Prefix" type:"string"`

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	RoleNameArn *string `locationName:"roleNameArn" min:"1" type:"string" required:"true"`

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	SnsTopicArn *string `locationName:"snsTopicArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GenerateDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GenerateDataSetInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateDataSetInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GenerateDataSetInput"}
	if s.CustomerDefinedValues != nil && len(s.CustomerDefinedValues) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CustomerDefinedValues", 1))
	}
	if s.DataSetPublicationDate == nil {
		invalidParams.Add(request.NewErrParamRequired("DataSetPublicationDate"))
	}
	if s.DataSetType == nil {
		invalidParams.Add(request.NewErrParamRequired("DataSetType"))
	}
	if s.DataSetType != nil && len(*s.DataSetType) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("DataSetType", 1))
	}
	if s.DestinationS3BucketName == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationS3BucketName"))
	}
	if s.DestinationS3BucketName != nil && len(*s.DestinationS3BucketName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("DestinationS3BucketName", 1))
	}
	if s.RoleNameArn == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleNameArn"))
	}
	if s.RoleNameArn != nil && len(*s.RoleNameArn) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RoleNameArn", 1))
	}
	if s.SnsTopicArn == nil {
		invalidParams.Add(request.NewErrParamRequired("SnsTopicArn"))
	}
	if s.SnsTopicArn != nil && len(*s.SnsTopicArn) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SnsTopicArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the result of the GenerateDataSet operation.
type GenerateDataSetOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier representing a specific request to the GenerateDataSet
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string `locationName:"dataSetRequestId" type:"string"`
}

// String returns the string representation
func (s GenerateDataSetOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GenerateDataSetOutput) GoString() string {
	return s.String()
}

// Container for the parameters to the StartSupportDataExport operation.
type StartSupportDataExportInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Key-value pairs which will be returned, unmodified, in the Amazon
	// SNS notification message and the data set metadata file.
	CustomerDefinedValues map[string]*string `locationName:"customerDefinedValues" min:"1" type:"map"`

	// Specifies the data set type to be written to the output csv file. The data
	// set types customer_support_contacts_data and test_customer_support_contacts_data
	// both result in a csv file containing the following fields: Product Id, Customer
	// Guid, Subscription Guid, Subscription Start Date, Organization, AWS Account
	// Id, Given Name, Surname, Telephone Number, Email, Title, Country Code, ZIP
	// Code, Operation Type, and Operation Time. Currently, only the test_customer_support_contacts_data
	// value is supported
	//
	//   customer_support_contacts_data Customer support contact data. The data
	// set will contain all changes (Creates, Updates, and Deletes) to customer
	// support contact data from the date specified in the from_date parameter.
	// test_customer_support_contacts_data An example data set containing static
	// test data in the same format as customer_support_contacts_data
	DataSetType *string `locationName:"dataSetType" min:"1" type:"string" required:"true" enum:"SupportDataSetType"`

	// The name (friendly name, not ARN) of the destination S3 bucket.
	DestinationS3BucketName *string `locationName:"destinationS3BucketName" min:"1" type:"string" required:"true"`

	// (Optional) The desired S3 prefix for the published data set, similar to a
	// directory path in standard file systems. For example, if given the bucket
	// name "mybucket" and the prefix "myprefix/mydatasets", the output file "outputfile"
	// would be published to "s3://mybucket/myprefix/mydatasets/outputfile". If
	// the prefix directory structure does not exist, it will be created. If no
	// prefix is provided, the data set will be published to the S3 bucket root.
	DestinationS3Prefix *string `locationName:"destinationS3Prefix" type:"string"`

	// The start date from which to retrieve the data set. This parameter only affects
	// the customer_support_contacts_data data set type.
	FromDate *time.Time `locationName:"fromDate" type:"timestamp" timestampFormat:"unix" required:"true"`

	// The Amazon Resource Name (ARN) of the Role with an attached permissions policy
	// to interact with the provided AWS services.
	RoleNameArn *string `locationName:"roleNameArn" min:"1" type:"string" required:"true"`

	// Amazon Resource Name (ARN) for the SNS Topic that will be notified when the
	// data set has been published or if an error has occurred.
	SnsTopicArn *string `locationName:"snsTopicArn" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StartSupportDataExportInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartSupportDataExportInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSupportDataExportInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "StartSupportDataExportInput"}
	if s.CustomerDefinedValues != nil && len(s.CustomerDefinedValues) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CustomerDefinedValues", 1))
	}
	if s.DataSetType == nil {
		invalidParams.Add(request.NewErrParamRequired("DataSetType"))
	}
	if s.DataSetType != nil && len(*s.DataSetType) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("DataSetType", 1))
	}
	if s.DestinationS3BucketName == nil {
		invalidParams.Add(request.NewErrParamRequired("DestinationS3BucketName"))
	}
	if s.DestinationS3BucketName != nil && len(*s.DestinationS3BucketName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("DestinationS3BucketName", 1))
	}
	if s.FromDate == nil {
		invalidParams.Add(request.NewErrParamRequired("FromDate"))
	}
	if s.RoleNameArn == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleNameArn"))
	}
	if s.RoleNameArn != nil && len(*s.RoleNameArn) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RoleNameArn", 1))
	}
	if s.SnsTopicArn == nil {
		invalidParams.Add(request.NewErrParamRequired("SnsTopicArn"))
	}
	if s.SnsTopicArn != nil && len(*s.SnsTopicArn) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SnsTopicArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Container for the result of the StartSupportDataExport operation.
type StartSupportDataExportOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier representing a specific request to the StartSupportDataExport
	// operation. This identifier can be used to correlate a request with notifications
	// from the SNS topic.
	DataSetRequestId *string `locationName:"dataSetRequestId" type:"string"`
}

// String returns the string representation
func (s StartSupportDataExportOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s StartSupportDataExportOutput) GoString() string {
	return s.String()
}

const (
	// @enum DataSetType
	DataSetTypeCustomerSubscriberHourlyMonthlySubscriptions = "customer_subscriber_hourly_monthly_subscriptions"
	// @enum DataSetType
	DataSetTypeCustomerSubscriberAnnualSubscriptions = "customer_subscriber_annual_subscriptions"
	// @enum DataSetType
	DataSetTypeDailyBusinessUsageByInstanceType = "daily_business_usage_by_instance_type"
	// @enum DataSetType
	DataSetTypeDailyBusinessFees = "daily_business_fees"
	// @enum DataSetType
	DataSetTypeDailyBusinessFreeTrialConversions = "daily_business_free_trial_conversions"
	// @enum DataSetType
	DataSetTypeDailyBusinessNewInstances = "daily_business_new_instances"
	// @enum DataSetType
	DataSetTypeDailyBusinessNewProductSubscribers = "daily_business_new_product_subscribers"
	// @enum DataSetType
	DataSetTypeDailyBusinessCanceledProductSubscribers = "daily_business_canceled_product_subscribers"
	// @enum DataSetType
	DataSetTypeMonthlyRevenueBillingAndRevenueData = "monthly_revenue_billing_and_revenue_data"
	// @enum DataSetType
	DataSetTypeMonthlyRevenueAnnualSubscriptions = "monthly_revenue_annual_subscriptions"
	// @enum DataSetType
	DataSetTypeDisbursedAmountByProduct = "disbursed_amount_by_product"
	// @enum DataSetType
	DataSetTypeDisbursedAmountByProductWithUncollectedFunds = "disbursed_amount_by_product_with_uncollected_funds"
	// @enum DataSetType
	DataSetTypeDisbursedAmountByCustomerGeo = "disbursed_amount_by_customer_geo"
	// @enum DataSetType
	DataSetTypeDisbursedAmountByAgeOfUncollectedFunds = "disbursed_amount_by_age_of_uncollected_funds"
	// @enum DataSetType
	DataSetTypeDisbursedAmountByAgeOfDisbursedFunds = "disbursed_amount_by_age_of_disbursed_funds"
	// @enum DataSetType
	DataSetTypeCustomerProfileByIndustry = "customer_profile_by_industry"
	// @enum DataSetType
	DataSetTypeCustomerProfileByRevenue = "customer_profile_by_revenue"
	// @enum DataSetType
	DataSetTypeCustomerProfileByGeography = "customer_profile_by_geography"
)

const (
	// @enum SupportDataSetType
	SupportDataSetTypeCustomerSupportContactsData = "customer_support_contacts_data"
	// @enum SupportDataSetType
	SupportDataSetTypeTestCustomerSupportContactsData = "test_customer_support_contacts_data"
)
