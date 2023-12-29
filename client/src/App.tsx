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
          onClick: hiddeModal
        }]
      })
    }
  }, [error])

  const showModal = (modalData: ModalData) => {
    setModalData(modalData)

    setModalOpen(true)
  }

  const hiddeModal = () => setModalOpen(false)

  return (
    <>
      {isModalOpen ? 
        <div className={styles.containerModal}>
          {isModalOpen && <Modal data={modalData} onClose={hiddeModal}/>}
        </div> :
        <div className={styles.containerMain}>
          <Routes>
            <Route path="/" element={<Companies />}/>
          </Routes>
      </div>
      }
    </>
  )
}

export default App
