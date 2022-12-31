export const initDom = (app) => {
  if (import.meta.env.MODE === 'development') {
    document.onmousedown = function(e) {
      console.log("initDom.e", e)
      if (e.shiftKey && e.button === 0) {
        e.preventDefault()
        sendRequestToOpenFileInEditor(getFilePath(e))
      }
    }
  } else{
    app.config.productionTip = false
  }
}

const getFilePath = (e) => {
  let element = e
  if (e.target) {
    element = e.target
  }
  if (!element || !element.getAttribute) return null
  if (element.getAttribute('code-location')) {
    return element.getAttribute('code-location')
  }
  return getFilePath(element.parentNode)
}

// 在编辑器中发送打开文件的请求
const sendRequestToOpenFileInEditor = (filePath) => {
  console.log("sendRequestToOpenFileInEditor.filePath", filePath)
  const protocol = window.location.protocol
    ? window.location.protocol
    : 'http:'
  const hostname = window.location.hostname
    ? window.location.hostname
    : 'localhost'
  const port = window.location.port ? window.location.port : '80'
  fetch(`${protocol}//${hostname}:${port}/gvaPositionCode?filePath=${filePath}`)
    .catch((error) => {
      console.log(error)
    })
}
