import Loader from "../loader"
import SelectableTable from "../table"
import Form from "../form"

import { ReactNode } from "react"
import { Row, SelectableTableProps } from "../table/interfaces"
import { FormProps } from "../form/interfaces"


import styles from './updatableTableWithFilters.module.scss'

interface UpdatableTableWithFiltersProps<T extends Row, F> {
    className?: string
    isLoading: boolean
    filters?: FormProps<F>
    table: SelectableTableProps<T>
    form: FormProps<F>
    children?: ReactNode
}

const UpdatableTableWithFilters: React.FC<UpdatableTableWithFiltersProps<any, any>> = ({
    className="",
    isLoading,
    filters,
    table,
    form,
    children
}) => {
    return (
        <div className={`${styles.containerMain} ${className}`}>
            <div className={styles.containerFilter}>
                {filters && <Form
                    title={filters.title}
                    fields={filters.fields}
                    buttons={filters.buttons}
                    formValues={filters.formValues}
                    handleChange={filters.handleChange}
                    onSubmit={filters.onSubmit}
                    onReset={filters.onReset}
                    refIndex={filters.refIndex}
                />}
            </div>
            <div className={styles.containerTable}>
                <Loader isLoading={isLoading}>
                    <SelectableTable
                        headers={table.headers}
                        rows={table.rows}
                        handleDoubleClick={table.handleDoubleClick}
                        selectedRow={form.formValues.id}
                    />
                </Loader>
            </div>
            <Form
                title={form.title}
                fields={form.fields}
                buttons={form.buttons}
                formValues={form.formValues}
                handleChange={form.handleChange}
                onSubmit={form.onSubmit}
                onReset={form.onReset}
                refIndex={form.refIndex}
            />
            {children}
        </div>
    )
}

export default UpdatableTableWithFilters
