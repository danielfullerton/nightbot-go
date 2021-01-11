import React, { Component } from 'react';
import { Channel } from '../types/Channel';

interface ChannelState {
  channel: Channel;
  lock: boolean;
}

export class ChannelComponent extends Component<any, ChannelState> {
  // constructor(props: any) {
  //   super(props);
  //   this.state = {
  //     channel: null,
  //     lock: false
  //   };
  // }
  //
  // async fetchChannel(): Promise<void> {
  //   if (!this.state.lock) {
  //     this.setState({ lock: true });
  //     const resp = await axios.get<ChannelResponse>('/api/channel');
  //     const { channel } = resp.data;
  //     this.setState({ channel: channel, lock: false });
  //   }
  // }
  //
  // componentDidMount() {
  //   setInterval(async () => {
  //     await this.fetchChannel();
  //   }, 3000);
  // }

  render() {
    return (
      <div>
        Hello world
      </div>
    );
  }
}
