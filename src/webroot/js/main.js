define(
	"main",
	[
		"MessageList"
	],
	function(MessageList) {
		var ws = new WebSocket("wss://ragtime-mypianoplayer.c9users.io/entry");
		var list = new MessageList(ws);
		ko.applyBindings(list);
	}
);
