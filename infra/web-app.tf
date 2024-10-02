resource "digitalocean_droplet" "web-app" {
  image = "ubuntu-20-04-x64"
  name = "web-app"
  region = "sgp1"
  size = "s-1vcpu-2gb"
  ssh_keys = [ data.digitalocean_ssh_key.terraform.id ]
  connection {
    host = self.ipv4_address
    user = "root"
    type = "ssh"
    private_key = file(var.pvt_key)
    timeout = "2m"
  }
  provisioner "remote-exec" {
    inline = [
      # install docker
      "sudo apt update",
      "sudo apt install -y docker.io",
      "sudo usermod -aG docker $USER",
      "sudo systemctl restart docker"
    ]
  }
}