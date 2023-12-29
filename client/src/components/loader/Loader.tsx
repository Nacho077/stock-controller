import { ReactNode } from 'react'
import styles from './loader.module.scss'

interface LoaderProps {
    children: ReactNode
    isLoading: boolean
}

const Loader: React.FC<LoaderProps> = ({children, isLoading}) => {
    if (isLoading) return <span className={styles.loader}></span>

    return children
}

export default Loader
