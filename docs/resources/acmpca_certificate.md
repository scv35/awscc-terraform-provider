---
page_title: "awscc_acmpca_certificate Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::ACMPCA::Certificate resource is used to issue a certificate using your private certificate authority. For more information, see the IssueCertificate https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html action.
---

# awscc_acmpca_certificate (Resource)

The ``AWS::ACMPCA::Certificate`` resource is used to issue a certificate using your private certificate authority. For more information, see the [IssueCertificate](https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html) action.

## Example Usage

### Basic example

```terraform
resource "awscc_acmpca_certificate" "example" {
  certificate_authority_arn   = awscc_acmpca_certificate_authority.example.arn
  certificate_signing_request = awscc_acmpca_certificate_authority.example.certificate_signing_request
  signing_algorithm           = "SHA256WITHRSA"

  template_arn = "arn:${data.aws_partition.current.partition}:acm-pca:::template/RootCACertificate/V1"

  validity = {
    type  = "YEARS"
    value = 5
  }
}

resource "awscc_acmpca_certificate_authority" "example" {
  key_algorithm     = "RSA_4096"
  signing_algorithm = "SHA512WITHRSA"
  type              = "ROOT"
  subject = {
    common_name = "example.com"
  }
}

resource "awscc_acmpca_certificate_authority_activation" "example" {
  certificate               = awscc_acmpca_certificate.example.certificate
  certificate_authority_arn = awscc_acmpca_certificate_authority.example.arn
}

data "aws_partition" "current" {}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `certificate_authority_arn` (String) The Amazon Resource Name (ARN) for the private CA issues the certificate.
- `certificate_signing_request` (String) The certificate signing request (CSR) for the certificate.
- `signing_algorithm` (String) The name of the algorithm that will be used to sign the certificate to be issued. 
 This parameter should not be confused with the ``SigningAlgorithm`` parameter used to sign a CSR in the ``CreateCertificateAuthority`` action.
  The specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.
- `validity` (Attributes) The period of time during which the certificate will be valid. (see [below for nested schema](#nestedatt--validity))

### Optional

- `api_passthrough` (Attributes) Specifies X.509 certificate information to be included in the issued certificate. An ``APIPassthrough`` or ``APICSRPassthrough`` template variant must be selected, or else this parameter is ignored. (see [below for nested schema](#nestedatt--api_passthrough))
- `template_arn` (String) Specifies a custom configuration template to use when issuing a certificate. If this parameter is not provided, PCAshort defaults to the ``EndEntityCertificate/V1`` template. For more information about PCAshort templates, see [Using Templates](https://docs.aws.amazon.com/privateca/latest/userguide/UsingTemplates.html).
- `validity_not_before` (Attributes) Information describing the start of the validity period of the certificate. This parameter sets the ?Not Before" date for the certificate.
 By default, when issuing a certificate, PCAshort sets the "Not Before" date to the issuance time minus 60 minutes. This compensates for clock inconsistencies across computer systems. The ``ValidityNotBefore`` parameter can be used to customize the ?Not Before? value. 
 Unlike the ``Validity`` parameter, the ``ValidityNotBefore`` parameter is optional.
 The ``ValidityNotBefore`` value is expressed as an explicit date and time, using the ``Validity`` type value ``ABSOLUTE``. (see [below for nested schema](#nestedatt--validity_not_before))

### Read-Only

- `arn` (String)
- `certificate` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--validity"></a>
### Nested Schema for `validity`

Required:

- `type` (String) Specifies whether the ``Value`` parameter represents days, months, or years.
- `value` (Number) A long integer interpreted according to the value of ``Type``, below.


<a id="nestedatt--api_passthrough"></a>
### Nested Schema for `api_passthrough`

Optional:

- `extensions` (Attributes) Specifies X.509 extension information for a certificate. (see [below for nested schema](#nestedatt--api_passthrough--extensions))
- `subject` (Attributes) Contains information about the certificate subject. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. (see [below for nested schema](#nestedatt--api_passthrough--subject))

<a id="nestedatt--api_passthrough--extensions"></a>
### Nested Schema for `api_passthrough.extensions`

Optional:

- `certificate_policies` (Attributes List) Contains a sequence of one or more policy information terms, each of which consists of an object identifier (OID) and optional qualifiers. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).
 In an end-entity certificate, these terms indicate the policy under which the certificate was issued and the purposes for which it may be used. In a CA certificate, these terms limit the set of policies for certification paths that include this certificate. (see [below for nested schema](#nestedatt--api_passthrough--extensions--certificate_policies))
- `custom_extensions` (Attributes List) Contains a sequence of one or more X.509 extensions, each of which consists of an object identifier (OID), a base64-encoded value, and the critical flag. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29) (see [below for nested schema](#nestedatt--api_passthrough--extensions--custom_extensions))
- `extended_key_usage` (Attributes List) Specifies additional purposes for which the certified public key may be used other than basic purposes indicated in the ``KeyUsage`` extension. (see [below for nested schema](#nestedatt--api_passthrough--extensions--extended_key_usage))
- `key_usage` (Attributes) Defines one or more purposes for which the key contained in the certificate can be used. Default value for each option is false. (see [below for nested schema](#nestedatt--api_passthrough--extensions--key_usage))
- `subject_alternative_names` (Attributes List) The subject alternative name extension allows identities to be bound to the subject of the certificate. These identities may be included in addition to or in place of the identity in the subject field of the certificate. (see [below for nested schema](#nestedatt--api_passthrough--extensions--subject_alternative_names))

<a id="nestedatt--api_passthrough--extensions--certificate_policies"></a>
### Nested Schema for `api_passthrough.extensions.certificate_policies`

Optional:

- `cert_policy_id` (String) Specifies the object identifier (OID) of the certificate policy under which the certificate was issued. For more information, see NIST's definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).
- `policy_qualifiers` (Attributes List) Modifies the given ``CertPolicyId`` with a qualifier. AWS Private CA supports the certification practice statement (CPS) qualifier. (see [below for nested schema](#nestedatt--api_passthrough--extensions--certificate_policies--policy_qualifiers))

<a id="nestedatt--api_passthrough--extensions--certificate_policies--policy_qualifiers"></a>
### Nested Schema for `api_passthrough.extensions.certificate_policies.policy_qualifiers`

Optional:

- `policy_qualifier_id` (String) Identifies the qualifier modifying a ``CertPolicyId``.
- `qualifier` (Attributes) Defines the qualifier type. AWS Private CA supports the use of a URI for a CPS qualifier in this field. (see [below for nested schema](#nestedatt--api_passthrough--extensions--certificate_policies--policy_qualifiers--qualifier))

<a id="nestedatt--api_passthrough--extensions--certificate_policies--policy_qualifiers--qualifier"></a>
### Nested Schema for `api_passthrough.extensions.certificate_policies.policy_qualifiers.qualifier`

Optional:

- `cps_uri` (String) Contains a pointer to a certification practice statement (CPS) published by the CA.




<a id="nestedatt--api_passthrough--extensions--custom_extensions"></a>
### Nested Schema for `api_passthrough.extensions.custom_extensions`

Optional:

- `critical` (Boolean) Specifies the critical flag of the X.509 extension.
- `object_identifier` (String) Specifies the object identifier (OID) of the X.509 extension. For more information, see the [Global OID reference database.](https://docs.aws.amazon.com/https://oidref.com/2.5.29)
- `value` (String) Specifies the base64-encoded value of the X.509 extension.


<a id="nestedatt--api_passthrough--extensions--extended_key_usage"></a>
### Nested Schema for `api_passthrough.extensions.extended_key_usage`

Optional:

- `extended_key_usage_object_identifier` (String) Specifies a custom ``ExtendedKeyUsage`` with an object identifier (OID).
- `extended_key_usage_type` (String) Specifies a standard ``ExtendedKeyUsage`` as defined as in [RFC 5280](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc5280#section-4.2.1.12).


<a id="nestedatt--api_passthrough--extensions--key_usage"></a>
### Nested Schema for `api_passthrough.extensions.key_usage`

Optional:

- `crl_sign` (Boolean) Key can be used to sign CRLs.
- `data_encipherment` (Boolean) Key can be used to decipher data.
- `decipher_only` (Boolean) Key can be used only to decipher data.
- `digital_signature` (Boolean) Key can be used for digital signing.
- `encipher_only` (Boolean) Key can be used only to encipher data.
- `key_agreement` (Boolean) Key can be used in a key-agreement protocol.
- `key_cert_sign` (Boolean) Key can be used to sign certificates.
- `key_encipherment` (Boolean) Key can be used to encipher data.
- `non_repudiation` (Boolean) Key can be used for non-repudiation.


<a id="nestedatt--api_passthrough--extensions--subject_alternative_names"></a>
### Nested Schema for `api_passthrough.extensions.subject_alternative_names`

Optional:

- `directory_name` (Attributes) Contains information about the certificate subject. The certificate can be one issued by your private certificate authority (CA) or it can be your private CA certificate. The Subject field in the certificate identifies the entity that owns or controls the public key in the certificate. The entity can be a user, computer, device, or service. The Subject must contain an X.500 distinguished name (DN). A DN is a sequence of relative distinguished names (RDNs). The RDNs are separated by commas in the certificate. The DN must be unique for each entity, but your private CA can issue more than one certificate with the same DN to the same entity. (see [below for nested schema](#nestedatt--api_passthrough--extensions--subject_alternative_names--directory_name))
- `dns_name` (String) Represents ``GeneralName`` as a DNS name.
- `edi_party_name` (Attributes) Represents ``GeneralName`` as an ``EdiPartyName`` object. (see [below for nested schema](#nestedatt--api_passthrough--extensions--subject_alternative_names--edi_party_name))
- `ip_address` (String) Represents ``GeneralName`` as an IPv4 or IPv6 address.
- `other_name` (Attributes) Represents ``GeneralName`` using an ``OtherName`` object. (see [below for nested schema](#nestedatt--api_passthrough--extensions--subject_alternative_names--other_name))
- `registered_id` (String) Represents ``GeneralName`` as an object identifier (OID).
- `rfc_822_name` (String) Represents ``GeneralName`` as an [RFC 822](https://docs.aws.amazon.com/https://datatracker.ietf.org/doc/html/rfc822) email address.
- `uniform_resource_identifier` (String) Represents ``GeneralName`` as a URI.

<a id="nestedatt--api_passthrough--extensions--subject_alternative_names--directory_name"></a>
### Nested Schema for `api_passthrough.extensions.subject_alternative_names.directory_name`

Optional:

- `common_name` (String) For CA and end-entity certificates in a private PKI, the common name (CN) can be any string within the length limit.
 Note: In publicly trusted certificates, the common name must be a fully qualified domain name (FQDN) associated with the certificate subject.
- `country` (String) Two-digit code that specifies the country in which the certificate subject located.
- `custom_attributes` (Attributes List) Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST?s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).
  Custom attributes cannot be used in combination with standard attributes. (see [below for nested schema](#nestedatt--api_passthrough--extensions--subject_alternative_names--directory_name--custom_attributes))
- `distinguished_name_qualifier` (String) Disambiguating information for the certificate subject.
- `generation_qualifier` (String) Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.
- `given_name` (String) First name.
- `initials` (String) Concatenation that typically contains the first letter of the *GivenName*, the first letter of the middle name if one exists, and the first letter of the *Surname*.
- `locality` (String) The locality (such as a city or town) in which the certificate subject is located.
- `organization` (String) Legal name of the organization with which the certificate subject is affiliated.
- `organizational_unit` (String) A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
- `pseudonym` (String) Typically a shortened version of a longer *GivenName*. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
- `serial_number` (String) The certificate serial number.
- `state` (String) State in which the subject of the certificate is located.
- `surname` (String) Family name. In the US and the UK, for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
- `title` (String) A title such as Mr. or Ms., which is pre-pended to the name to refer formally to the certificate subject.

<a id="nestedatt--api_passthrough--extensions--subject_alternative_names--directory_name--custom_attributes"></a>
### Nested Schema for `api_passthrough.extensions.subject_alternative_names.directory_name.custom_attributes`

Optional:

- `object_identifier` (String) Specifies the object identifier (OID) of the attribute type of the relative distinguished name (RDN).
- `value` (String) Specifies the attribute value of relative distinguished name (RDN).



<a id="nestedatt--api_passthrough--extensions--subject_alternative_names--edi_party_name"></a>
### Nested Schema for `api_passthrough.extensions.subject_alternative_names.edi_party_name`

Optional:

- `name_assigner` (String) Specifies the name assigner.
- `party_name` (String) Specifies the party name.


<a id="nestedatt--api_passthrough--extensions--subject_alternative_names--other_name"></a>
### Nested Schema for `api_passthrough.extensions.subject_alternative_names.other_name`

Optional:

- `type_id` (String) Specifies an OID.
- `value` (String) Specifies an OID value.




<a id="nestedatt--api_passthrough--subject"></a>
### Nested Schema for `api_passthrough.subject`

Optional:

- `common_name` (String) For CA and end-entity certificates in a private PKI, the common name (CN) can be any string within the length limit.
 Note: In publicly trusted certificates, the common name must be a fully qualified domain name (FQDN) associated with the certificate subject.
- `country` (String) Two-digit code that specifies the country in which the certificate subject located.
- `custom_attributes` (Attributes List) Contains a sequence of one or more X.500 relative distinguished names (RDNs), each of which consists of an object identifier (OID) and a value. For more information, see NIST?s definition of [Object Identifier (OID)](https://docs.aws.amazon.com/https://csrc.nist.gov/glossary/term/Object_Identifier).
  Custom attributes cannot be used in combination with standard attributes. (see [below for nested schema](#nestedatt--api_passthrough--subject--custom_attributes))
- `distinguished_name_qualifier` (String) Disambiguating information for the certificate subject.
- `generation_qualifier` (String) Typically a qualifier appended to the name of an individual. Examples include Jr. for junior, Sr. for senior, and III for third.
- `given_name` (String) First name.
- `initials` (String) Concatenation that typically contains the first letter of the *GivenName*, the first letter of the middle name if one exists, and the first letter of the *Surname*.
- `locality` (String) The locality (such as a city or town) in which the certificate subject is located.
- `organization` (String) Legal name of the organization with which the certificate subject is affiliated.
- `organizational_unit` (String) A subdivision or unit of the organization (such as sales or finance) with which the certificate subject is affiliated.
- `pseudonym` (String) Typically a shortened version of a longer *GivenName*. For example, Jonathan is often shortened to John. Elizabeth is often shortened to Beth, Liz, or Eliza.
- `serial_number` (String) The certificate serial number.
- `state` (String) State in which the subject of the certificate is located.
- `surname` (String) Family name. In the US and the UK, for example, the surname of an individual is ordered last. In Asian cultures the surname is typically ordered first.
- `title` (String) A title such as Mr. or Ms., which is pre-pended to the name to refer formally to the certificate subject.

<a id="nestedatt--api_passthrough--subject--custom_attributes"></a>
### Nested Schema for `api_passthrough.subject.custom_attributes`

Optional:

- `object_identifier` (String) Specifies the object identifier (OID) of the attribute type of the relative distinguished name (RDN).
- `value` (String) Specifies the attribute value of relative distinguished name (RDN).




<a id="nestedatt--validity_not_before"></a>
### Nested Schema for `validity_not_before`

Optional:

- `type` (String) Specifies whether the ``Value`` parameter represents days, months, or years.
- `value` (Number) A long integer interpreted according to the value of ``Type``, below.

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, the [`import` block](https://developer.hashicorp.com/terraform/language/import) can be used with the `id` attribute, for example:

```terraform
import {
  to = awscc_acmpca_certificate.example
  id = "arn|certificate_authority_arn"
}
```

The [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import) can be used, for example:

```shell
$ terraform import awscc_acmpca_certificate.example "arn|certificate_authority_arn"
```
