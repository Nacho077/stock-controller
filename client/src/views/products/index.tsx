import { ChangeEvent, useState } from "react"
import { useParams } from "react-router-dom"

import UpdatableTableWithFilters from "../../components/updatableTableWithFilters"
import { Product } from "./interfaces"

const Products: React.FC = () => {
    const companyId = parseInt(useParams()["companyId"] || '0', 10)
    const rows: Product[] = []
    const [productForm, setProductForm] = useState({})

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target

        setProductForm({
            ...productForm,
            [name]: value
        })
    }

    return (
        <>
            <UpdatableTableWithFilters 
                isLoading={false}
                // filters={}
                table={{
                    headers: [],
                    rows: rows,
                }}
                form={{
                    title: "Nuevo Producto",
                    fields: [],
                    buttons: [{
                        title: "Limpiar",
                        type: "reset"
                    }, {
                        title: "Aceptar",
                        type: "submit"
                    }],
                    formValues: productForm,
                    handleChange: handleChange,
                    onSubmit: () => console.log("funciona")
                }}
            />
        </>
    )
}

export default Products
