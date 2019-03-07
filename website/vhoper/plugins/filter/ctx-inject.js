export default ({ app }, inject) => {
  // Set the function directly on the context.app object
  app.s2date = string => console.log('Okay, another function', string)
}
