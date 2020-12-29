function submitCredentials() {
  const clientId = document.getElementById("clientId").value;
  const clientSecret = document.getElementById("clientSecret").value;

  if (!clientId || !clientSecret) {
    return;
  }

  axios.post('/credentials', {
    clientId,
    clientSecret
  })
    .then(r => {
      const redirectUrl = r.data.redirectUrl;
      window.location.replace(redirectUrl);
    })
    .catch(console.error);
}
