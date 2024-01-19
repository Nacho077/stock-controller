import { useNavigate, useLocation } from 'react-router-dom'
import GoBackIcon from '../../assets/back.svg'
import styles from './navBar.module.scss'
import { useEffect, useState } from 'react'

const NavBar: React.FC = () => {
    const navigate = useNavigate()
    const location = useLocation()
    const [title, setTitle] = useState<string>("")

    const getTitle = (path: string): string => {
        switch(true){
            case /^\/company\/create/i.test(path):
                return "Crear una empresa"
            case /^\/company.*/i.test(path):
                return "Movimientos"
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
