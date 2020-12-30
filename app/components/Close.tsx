import React, { Component } from 'react';

export class Close extends Component<any, any> {
  render() {
    return (
      <div className="container credentials-page">
        <h2>Success! You may now close this window.</h2>
        <hr />
        <p>NightBot Go has successfully connected to your NightBot account and your widgets can now be used.</p>
      </div>
    );
  }
}
