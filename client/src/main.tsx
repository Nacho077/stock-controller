import ReactDOM from 'react-dom/client'

import { Provider } from 'react-redux'
import store from './configs/redux/store.ts'

import { BrowserRouter } from 'react-router-dom'
import App from './App.tsx'

import './index.scss'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <Provider store={store}>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </Provider>,
)
