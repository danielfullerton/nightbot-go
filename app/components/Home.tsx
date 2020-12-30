import React, { Component } from 'react';

export class Home extends Component<any, any> {
  submitCredentials(skipCredentials?: boolean) {
    console.log(skipCredentials);
  }

  render() {
    return (
      <div>
        <div className="container credentials-page">
          <h2>Enter NightBot API Credentials</h2>
          <hr/>
          <div className="row credentials-form">
            <form>
              <div className="form-group">
                <label htmlFor="clientId">Client ID:</label>
                <input type="text" className="form-control" id="clientId" aria-describedby="clientId"
                       placeholder="Enter NightBot Client ID"/>
              </div>
              <div className="form-group">
                <label htmlFor="clientSecret">Client Secret:</label>
                <input type="password" className="form-control" id="clientSecret" aria-describedby="clientSecret"
                       placeholder="Enter NightBot Client Secret" autoComplete=""/>
              </div>
              <button type="button" className="btn btn-primary" onClick={() => this.submitCredentials(true)}>Submit</button>
              <button className="btn btn-success" onClick={() => this.submitCredentials(true)}>Use Environment Variables
              </button>
            </form>
          </div>

          <h2>Instructions</h2>
          <hr/>
          <div className="row credentials-instructions">
            <i>if you entered your Client ID and Client Secret as environment variables when starting NightBot Go,
              please simply click "Use Environment Variables"</i>
            <br/>
            <ol>
              <li>Log in to the
                <a href="https://nightbot.tv/account/applications" target="_blank">NightBot Applications Page</a>
              </li>
              <li>Click the "New App" button.</li>
              <li>
                Enter <b>NightBot Go</b> for the name. <i>The name you use for the application is not important -
                you can use any value.</i>
              </li>
              <li>
                Enter your NightBot Go URL in the Redirect URIs box with the endpoint "/token". For example, if you
                are running NightBot Go on your local
                machine using port 5775, you would enter "http://localhost:5775/token".
                <img src="/assets/AddAnApplication.png" alt="Add An Application"/>
              </li>
              <li>Now click the pencil icon next to "NightBot Go" in the My Apps section, and click "New Secret" in
                the pop up box.
              </li>
              <li>
                Copy and paste the Client ID and Client Secret in the form above.
                <img src="/assets/EditApplication.png" alt="Edit Application"/>
              </li>
              <li>Finally, hit the submit button on this page. You will be redirected to NightBot to authorize this
                application to access
                some of your NightBot data, such as your queue information and channel information.
              </li>
            </ol>
          </div>
        </div>
      </div>
    );
  }
}
