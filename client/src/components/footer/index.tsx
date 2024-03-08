import GithubLogo from '../../assets/github.svg'
import LinkedInLogo from '../../assets/linkedin.svg'
import MailLogo from '../../assets/mail.svg'
import styles from './footer.module.scss'

const Footer: React.FC = () => {
    return (
    <footer className={styles.container}>
        <div className={styles.containerInfo}>
            <img src={GithubLogo} />
            <a target="_blank" href="https://github.com/LeyAylen6">LeyAylen6</a>
            <a target="_blank" href="https://github.com/Nacho077">Nacho077</a>
        </div>
        <div className={styles.containerInfo}>
            <img src={LinkedInLogo} />
            <a target="_blank" href="https://www.linkedin.com/in/leilaaylensalguero/">Leila Salguero</a>
            <a target="_blank" href="https://www.linkedin.com/in/ignacio-ezequiel-gimenez/">Ignacio Gimenez</a>
        </div>
        <div className={styles.containerInfo}>
            <img src={MailLogo} />
            <a target="_blank" href="mailto:leiisalguero@gmail.com">leiisalguero@gmail.com</a>
            <a target="_blank" href="mailto:ignaciogimenez70@gmail.com">ignaciogimenez70@gmail.com</a>
        </div>
    </footer>
    )
}

export default Footer
