import { ChangeEvent, useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { useAppSelector, useCreateNewMovement, useCreateNewProduct } from '../../hooks'

import UpdatableTableWithFilters from '../../components/updatableTableWithFilters'

import { movementsFiltersFields, ProductFilters, ProductMovement, getDefaultFilters, getDefaultMovement, movementFormFields, movementHeaders } from './interfaces'

import styles from './movements.module.scss'
import { Link } from 'react-router-dom'
import { useGetProductsByCompanyId, useGetProductsMovementsFiltered, useUpdateMovement } from '../../hooks'
import { Product } from '../products/interfaces'

const Movements: React.FC = () => {
    const companyId = parseInt(useParams()["companyId"] || '0', 10)
    const useFilters = useGetProductsMovementsFiltered()
    useGetProductsByCompanyId(companyId)
    const { movements: rows, totalUnits, products } = useAppSelector(state => state.reducer)
    const [filters, setFilters] = useState<ProductFilters>(getDefaultFilters())
    const [movementForm, setMovementForm] = useState<ProductMovement>(getDefaultMovement(companyId))
    const createNewMovement = useCreateNewMovement()
    const updateMovement = useUpdateMovement()
    const createNewProduct = useCreateNewProduct()
    const [isLoading, setLoading] = useState<boolean>(false)

    useEffect(() => {
        useFilters(companyId, filters)
    }, [])

    const updateMovements = (filters: ProductFilters) => {
        setLoading(true)
        useFilters(companyId, filters)
            .then(() => setLoading(false))
    }

    useEffect(() => updateMovements(filters), [])


    const autoCompleteFields = <T extends ProductFilters | ProductMovement>(state: T, codeToFind: string): T => {
        const movement = rows.find(row => row.code?.toLowerCase() == codeToFind?.toLowerCase())
        const productFounded = products.find(product => product.code?.toLowerCase() == codeToFind?.toLowerCase())

        return {
            ...state,
            brand: movement?.brand || productFounded?.brand || state.brand,
            name: movement?.name || productFounded?.name || state.name,
            detail: movement?.detail || (state as ProductMovement).detail,
            deposit: movement?.deposit || (state as ProductMovement).deposit,
            observations: movement?.observations || (state as ProductMovement).observations,
            productId: productFounded?.id,
        }
    }

    const handleFilters = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target

        let newState = {
            ...filters,
            [name]: value
        }

        if (name == "code") {
            newState = autoCompleteFields(newState, value)
        }

        setFilters(newState)
    }

    const clearFilters = () => {
        setFilters(getDefaultFilters())
        updateMovements(getDefaultFilters())
    }

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target

        let newState = {
            ...movementForm,
            [name]: value
        }

        if (name == "code") {
            newState = autoCompleteFields(newState, value)
        }

        setMovementForm(newState)
    }

    const handleDoubleClick = (movement: ProductMovement) => {
        setMovementForm(movement)
    }

    const resetForm = () => {
        setMovementForm(getDefaultMovement(companyId))
    }

    const handleSubmit = async () => {
        let request = movementForm

        if (movementForm.movementId === 0) {
            if (!movementForm.productId && movementForm.code !== "") {
                const result = await createNewProduct({
                    id: 0,
                    name: movementForm.name,
                    code: movementForm.code,
                    brand: movementForm.brand,
                    detail: movementForm.detail,
                    company_id: companyId,
                })

                if ('data' in result) {
                    request = {
                        ...movementForm,
                        productId: result.data.id
                    }
                }
            }

            createNewMovement(companyId, request)
        } else {
            updateMovement(companyId, movementForm.movementId, movementForm)
        }

        setMovementForm(getDefaultMovement(
            companyId,
            request.date,
            request.shippingCode,
            request.code,
            request.productId
        ))
    }

    const getAvailableProductsToAutocomplete = <T extends ProductFilters | ProductMovement>(form: T): Product[] => {
        if (form.code === "" && form.name === "" && form.brand === "") return products

        return products.filter(product =>
            (form.code !== "" && product.code.includes(movementForm.code)) ||
            (form.name !== "" && product.name.includes(movementForm.name)) ||
            (form.brand !== "" && product.brand.includes(movementForm.brand))
        )
    }

    return (
        <UpdatableTableWithFilters
            className={styles.containerMain}
            isLoading={isLoading}
            filters={{
                fields: movementsFiltersFields(getAvailableProductsToAutocomplete(filters)),
                buttons: [{
                    title: "Vaciar Filtros",
                    type: "reset"
                }, {
                    title: "Buscar",
                    type: "submit"
                }],
                formValues: filters,
                handleChange: handleFilters,
                onSubmit: () => updateMovements(filters),
                onReset: clearFilters,
                refIndex: -1
            }}
            table={{
                headers: movementHeaders,
                rows: rows,
                handleDoubleClick: handleDoubleClick
            }}
            form={{
                title: movementForm.movementId === 0 ? "Nuevo Movimiento" : "Modificar Movimiento",
                fields: movementFormFields(movementForm.movementId !== 0, getAvailableProductsToAutocomplete(movementForm)),
                buttons: [{
                    title: "Limpiar",
                    type: "reset"
                }, {
                    title: "Aceptar",
                    type: "submit"
                }],
                formValues: movementForm,
                handleChange: handleChange,
                onSubmit: handleSubmit,
                onReset: resetForm,
                refIndex: 2
            }}
        >
            <div className={styles.containerInfo}>
                <div>
                    <span>Cant Total</span>
                    <span>{totalUnits}</span>
                </div>
                <Link to={`/company/${companyId}/products`} className={styles.button}>ver productos</Link>
            </div>
        </UpdatableTableWithFilters>
    )
}

export default Movements
