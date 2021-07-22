const urlImage=(url)=> {
    if (!url) {
      return ""
    }
    return `http://localhost:3000/images/${url}`
  }
  export default urlImage