const copy = function(obj) {
  const newobj = {}
  for (const attr in obj) {
    newobj[attr] = obj[attr]
  }
  return newobj
}

function writeObj(obj) {
  let description = ''
  for (const i in obj) {
    const property = obj[i]
    description += i + ' = ' + property + '\n'
  }
  console.log('obj:' + description)
}
export { copy, writeObj }
