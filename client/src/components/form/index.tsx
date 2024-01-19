import { FormEvent, useRef } from "react"
import { FormProps } from "./interfaces"

import styles from './form.module.scss'

const Form = <T extends Record<string, string | number>>({
    title,
    fields,
    buttons,
    formValues,
    handleChange,
    onSubmit,
    onReset,
    refIndex=-1
}: FormProps<T>) => {
    const inputRef = useRef<HTMLInputElement>(null)
    
    const handleSubmit = (e:FormEvent) => {
        e.preventDefault()

        onSubmit(e)
        inputRef.current?.focus()
    }

    return <form onSubmit={handleSubmit} className={styles.containerForm} onReset={onReset}>
        {title && <h3>{title}</h3>}
        <div className={styles.containerInputs}>
            {fields.map(({name, label, datalist, ...opts}, i) => (
                <div className={styles.containerInput} key={name}>
                    <label htmlFor={`${name}-input`}>{label}</label>
                    <input 
                        id={`${name}-input`}
                        name={name}
                        value={formValues[name]}
                        onChange={handleChange}
                        list={name}
                        ref={i == refIndex ? inputRef : undefined}
                        {...opts}
                    />
                    {datalist.length > 0 && 
                        <datalist id={name}>
                            {datalist.map(opt => <option value={opt} key={opt} />)}
                        </datalist>
                    }
                </div>
            ))}
        </div>
        <div className={styles.containerButtons}>
            {buttons.map(({title, type="button", onClick=() => {}}) => <button
                key={title}
                className={styles[type]}
                type={type}
                onClick={onClick}
            >{title}</button>)}
        </div>
    </form>
}

export default Form
