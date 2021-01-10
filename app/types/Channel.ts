export interface Channel {
  displayName: string;
  joined: boolean;
  name: string;
  _id: string;
}

export interface ChannelResponse {
  channel: Channel;
  status: number;
}
