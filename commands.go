package limitless

//On turns group on, or all
// List of group: 1, 2, 3, 4, all
func (c *Client) On(group string) {
	c.command("group_" + group + "_on")
}

func (c *Client) Off(group string) {
	c.command("group_" + group + "_off")
}

func (c *Client) SetColor(group string, color string) {
	c.On(group)
	c.color(color)
}

func (c *Client) color(color string) {
	c.udpClient.Write([]byte(Commands["color"] + Colors[color] + "\x55"))
}

func (c *Client) command(cmd string) {
	c.udpClient.Write([]byte(Commands[cmd] + "\x00\x55"))
}
