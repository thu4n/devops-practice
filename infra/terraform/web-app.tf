resource "digitalocean_droplet" "web-app" {
  image    = "ubuntu-22-04-x64"
  name     = "web-app"
  region   = "sgp1"
  size     = "s-2vcpu-2gb"
  ssh_keys = [data.digitalocean_ssh_key.bgsv-ubuntu.id]
  connection {
    host        = self.ipv4_address
    user        = "root"
    type        = "ssh"
    private_key = file(var.pvt_key)
    timeout     = "2m"
  }
  provisioner "remote-exec" {
    inline = ["echo Remote execution is successful!"]
  }
  provisioner "local-exec" {
    command = "ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u root -i '${self.ipv4_address},' --private-key ${var.pvt_key} -e 'pub_key=${var.pub_key}' ../ansible/install_docker.yml"
  }
  provisioner "local-exec" {
    command = "ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u root -i '${self.ipv4_address},' --private-key ${var.pvt_key} -e 'pub_key=${var.pub_key}' ../ansible/install_webapp.yml"
  }
}

resource "digitalocean_firewall" "web-app-firewall" {
  name        = "web-app-firewall"
  droplet_ids = [digitalocean_droplet.web-app.id]

  # Web UI
  inbound_rule {
    protocol         = "tcp"
    port_range       = "80"
    source_addresses = ["0.0.0.0/0"]
  }

  # SSH
  inbound_rule {
    protocol         = "tcp"
    port_range       = "22"
    source_addresses = ["0.0.0.0/0"]
  }
}