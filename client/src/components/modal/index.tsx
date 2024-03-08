import { ModalInfo } from './interfaces'
import styles from './modal.module.scss'

const Modal: React.FC<ModalInfo> = ({data: {text, buttons}, onClose}) => {
    return (
        <div className={styles.containerModal}>
            <div className={styles.modal}>
                <button className={styles.containerClose} onClick={onClose}>
                    X
                </button>
                <div className={styles.containerText}>
                    <p>{text}</p>
                </div>
                <div className={styles.containerButtons}>
                    {buttons.map(button => (
                        <button
                        key={button.text}
                        className={styles.button}
                        onClick={() => {
                            button.onClick && button.onClick()
                            onClose()
                        }}
                        >
                            {button.text}
                        </button>
                    ))}
                </div>
            </div>
        </div>
    )
}

export default Modal
