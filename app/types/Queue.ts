export interface Track {
  providerId: string;
  provider: string;
  duration: number;
  title: string;
  artist: string;
  url: string;
}

export interface User {
  name: string;
  displayName: string;
  providerId: string;
  provider: string;
}

export interface Song {
  _id: string;
  createdAt: string;
  updatedAt: string;
  track: Track;
  user: User;
  _position: number;
}

export type Queue = Song[];

export interface QueueResponse {
  _total: number;
  _currentSong: Song;
  _requestsEnabled: boolean;
  _providers: string[];
  _searchProvider: string;
  status: number;
  queue: Queue;
}
