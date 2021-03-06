package packethandler

import "github.com/bnch/bancho/packets"

// ChatMessage is a message in the osu! chat.
type ChatMessage struct {
	From    string
	To      string
	Content string
	UserID  int32
}

// ToPacket converts a ChatMessage to a packets.Packet.
func (c ChatMessage) ToPacket(s *Session) packets.Packet {
	pack := c.ToPacketNoIgnore()
	pack.Ignored = append(pack.Ignored, s.User.Token)
	return pack
}

// ToPacketNoIgnore returns a chat message without the ignored sender. Used for PMs.
func (c ChatMessage) ToPacketNoIgnore() packets.Packet {
	return packets.ChatMessage(c.From, c.To, c.Content, c.UserID)
}

// HandleMessage broadcasts a received message to all users.
func HandleMessage(ps packSess) {
	m := handleMessageInitial(ps)

	st := GetStream("chan/" + m.To)
	if st == nil {
		return
	}
	if !st.IsSubscribed(ps.s.User.Token) {
		SendMessage(ps.s.User.Token, "You haven't joined that channel.")
	}
	st.Send(m.ToPacket(ps.s))
}

// HandlePrivateMessage handles a message sent to a specific user.
func HandlePrivateMessage(ps packSess) {
	m := handleMessageInitial(ps)

	if m.To == BotName {
		// TODO: implement calling of bot library.
		return
	}

	for _, v := range CopySessions() {
		// TODO: usernameToSession
		if v.User.Name == m.To {
			v.Push(m.ToPacketNoIgnore())
			return
		}
	}
}

func handleMessageInitial(ps packSess) ChatMessage {
	m := ChatMessage{}
	ps.p.Unmarshal(&m.From, &m.Content, &m.To, &m.UserID)

	m.From = ps.s.User.Name
	m.UserID = ps.s.User.ID

	return m
}
