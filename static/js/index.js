function submitCredentials(skipCredentials) {
  let req;

  if (skipCredentials) {
    req = axios.post('/credentials?skipCredentials=1')
  } else {
    const clientId = document.getElementById("clientId").value;
    const clientSecret = document.getElementById("clientSecret").value;

    if (!clientId || !clientSecret) {
      return;
    }

    req = axios.post('/credentials', {
      clientId,
      clientSecret
    });
  }

  req
    .then(r => {
      const redirectUrl = r.data.redirectUrl;
      window.location.replace(redirectUrl);
    })
    .catch(console.error);
}
