
resource "aci_epg_useg_sub_block_statement" "full_example_epg_useg_block_statement" {
  parent_dn   = aci_epg_useg_block_statement.example.id
  annotation  = "annotation"
  description = "description"
  match       = "all"
  name        = "sub_criterion"
  name_alias  = "name_alias"
  owner_key   = "owner_key"
  owner_tag   = "owner_tag"
  annotations = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
  tags = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
}

resource "aci_epg_useg_sub_block_statement" "full_example_epg_useg_sub_block_statement" {
  parent_dn   = aci_epg_useg_sub_block_statement.example.id
  annotation  = "annotation"
  description = "description"
  match       = "all"
  name        = "sub_criterion"
  name_alias  = "name_alias"
  owner_key   = "owner_key"
  owner_tag   = "owner_tag"
  annotations = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
  tags = [
    {
      key   = "key_0"
      value = "value_1"
    }
  ]
}