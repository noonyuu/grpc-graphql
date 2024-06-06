const sdk = require("node-appwrite");

let client = new sdk.Client();

client
  .setEndpoint("https://cloud.appwrite.io/v1") // Your API Endpoint
  .setProject("") // Your project ID
  .setKey("") // Your secret API key
  .setSelfSigned(); // Use only on dev mode with a self-signed SSL cert

let users = new sdk.Users(client);

try {
  let res = users.create(sdk.ID.unique(), "email@example.com", "+123456789", "password", "Walter O'Brien");
} catch (e) {
  console.log(e.message);
}

let promise = users.create(sdk.ID.unique(), "email@example.com", "+123456789", "password", "Walter O'Brien");

promise.then(
  function (response) {
    console.log(response);
  },
  function (error) {
    console.log(error);
  }
);
