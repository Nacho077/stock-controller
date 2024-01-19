import { Routes, Route } from 'react-router-dom'
import { useAppSelector } from './hooks'
import { useEffect, useState } from 'react'

import NavBar from './components/navBar'
import Modal from './components/modal'
import Companies from './views/companies'
import Movements from './views/movements'

import { ModalData } from './components/modal/interfaces'

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
        <NavBar />
        <main>
          <Routes>
            <Route path="/" element={<Companies />}/> {/* agregar footer */}
            <Route path="/company/create" element={<h1>TEST</h1>} />
            <Route path="/company/:companyId" element={<Movements />}/>
          </Routes>
        </main>
      </div>
    </>
  )
}

export default App
