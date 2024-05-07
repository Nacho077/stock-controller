import { ChangeEvent, useState } from "react"
import { useParams } from "react-router-dom"

import UpdatableTableWithFilters from "../../components/updatableTableWithFilters"
import { useCreateNewProduct, useUpdateProduct } from "../../hooks"
import { useAppSelector } from "../../hooks"
import { Product, getDefaultProduct, getProduct, productFormFields, productsHeaders } from "./interfaces"

import styles from './products.module.scss'

const Products: React.FC = () => {
    const companyId = parseInt(useParams()["companyId"] || '0', 10)
    const rows = useAppSelector(state => state.reducer.products)
    const [productForm, setProductForm] = useState<Product>(getDefaultProduct(companyId))
    const createNewProduct = useCreateNewProduct()
    const updateProduct = useUpdateProduct()

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target

        setProductForm({
            ...productForm,
            [name]: value
        })
    }

    const resetForm = () => setProductForm(getDefaultProduct(companyId))

    const handleDoubleClick = (product : Product) => setProductForm(getProduct(product, companyId))

    const handleSubmit = () => {
        if (productForm.id === 0) {
            createNewProduct(productForm)
        } else {
            updateProduct(productForm.id, productForm)
        }
        
        resetForm()
    }

    return (
        <>
            <UpdatableTableWithFilters 
                className={styles.containerMain}
                isLoading={false}
                table={{
                    headers: productsHeaders,
                    rows: rows,
                    handleDoubleClick: handleDoubleClick
                }}
                form={{
                    title: "Nuevo Producto",
                    fields: productFormFields,
                    buttons: [{
                        title: "Limpiar",
                        type: "reset"
                    }, {
                        title: "Aceptar",
                        type: "submit"
                    }],
                    formValues: productForm,
                    handleChange: handleChange,
                    onSubmit: handleSubmit,
                    onReset: resetForm,
                    refIndex: 0
                }}
            />
        </>
    )
}

export default Products
