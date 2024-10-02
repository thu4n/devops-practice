resource "digitalocean_droplet" "jenkins-server" {
  image = "ubuntu-20-04-x64"
  name = "jenkins-server"
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
      "sudo apt update",
      "sudo apt install -y default-jdk jenkins -y",
      "sudo systemctl enable jenkins",
      "sudo systemctl start jenkins"
    ]
  }
}