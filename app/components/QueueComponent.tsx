import React, { Component } from 'react';
import { Queue, QueueResponse, Song } from '../types/Queue';
import axios from 'axios';

export interface QueueState {
  lock: boolean;
  queue: Queue;
  currentSong: Song;
}

export class QueueComponent extends Component<any, QueueState> {
  constructor(props: any) {
    super(props);
    this.state = {
      lock: false,
      queue: [],
      currentSong: null
    }
  }

  async fetchQueue() {
    if (!this.state.lock) {
      this.setState({ lock: true });
      const queueResponse = await axios.get<QueueResponse>('/api/queue');
      const { _currentSong, queue } = queueResponse.data;
      this.setState({ lock: false, queue: queue, currentSong: _currentSong });
    }
  }

  componentDidMount() {
    setInterval(async () => {
      await this.fetchQueue();
    }, 3000);
  }

  render() {
    return (
      <div>
        {this.state.currentSong?.track.title}
      </div>
    );
  }
}
