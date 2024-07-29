resource "local_file" "Pet" {
  filename = var.filename
  content  = "My pet name is ${var.PetName} Pet ID is ${random_integer.pet_id.result}"
}

resource "random_integer" "pet_id" {
  min = 1
  max = 10
}
output "Dog_info" {
    value= "local_file.Pet.content"
}