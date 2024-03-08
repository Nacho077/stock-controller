import { useNavigate, useLocation } from 'react-router-dom'
import GoBackIcon from '../../assets/back.svg'
import styles from './navBar.module.scss'
import { useEffect, useState } from 'react'
import { useAppSelector } from '../../hooks'

const NavBar: React.FC = () => {
    const navigate = useNavigate()
    const location = useLocation()
    const actualCompany = useAppSelector(state => state.reducer.actualCompany)
    const [title, setTitle] = useState<string>("")

    const getTitle = (path: string): string => {
        switch(true){
            case /^\/company\/create/i.test(path):
                return "Crear una empresa"
            case /^\/company\/[0-9]+/i.test(path):
                return actualCompany
            default: 
                return "Bienvenido"
        }
    }

    useEffect(() => {
        setTitle(getTitle(location.pathname))
    }, [location.pathname])

    return (
        <nav className={styles.container}>
            <img src={GoBackIcon} onClick={() => navigate(-1)} alt="back"/>
            <h1>{title}</h1>
            <span></span>
        </nav>
    )
}

export default NavBar
