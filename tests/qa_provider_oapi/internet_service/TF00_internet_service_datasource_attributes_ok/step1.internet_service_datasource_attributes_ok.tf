resource "outscale_internet_service" "outscale_internet_service" {
  tags {
       key = "newTag"
       value = "IGW"
       }
}

data "outscale_internet_service" "outscale_internet_serviced" {
    filter {
        name   = "internet_service_ids"
        values = [outscale_internet_service.outscale_internet_service.internet_service_id]
    }
}    
