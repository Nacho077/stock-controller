import { Routes, Route } from 'react-router-dom'
import { useAppSelector } from './hooks'
import { useEffect, useState } from 'react'

import Companies from './components/companies/Companies'
import Modal from './components/modal/Modal'

import { ModalData } from './interfaces/interfaces'
import styles from './App.module.scss'

function App() {
  const [isModalOpen, setModalOpen] = useState<boolean>(false)
  const [modalData, setModalData] = useState<ModalData>({text: "", buttons: []})
  const error = useAppSelector(state => state.reducer.error)

  useEffect(() => {
    if (error !== "") {
      showModal({
        text: error,
        buttons: [{
          text: "Aceptar",
          onClick: hideModal
        }]
      })
    }
  }, [error])

  const showModal = (modalData: ModalData) => {
    setModalData(modalData)

    setModalOpen(true)
  }

  const hideModal = () => setModalOpen(false)

  return (
    <>
      <div className={styles.containerMain}>
      {isModalOpen && <div className={styles.modalContainer}><Modal data={modalData} onClose={hideModal}/></div>}
        <Routes>
          <Route path="/" element={<Companies />}/>
        </Routes>
      </div>
    </>
  )
}

export default App
