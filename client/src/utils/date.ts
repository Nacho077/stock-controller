export const getDateNowString = () => {
    const date = new Date().toLocaleDateString()

    const [day, month, year] = date.split("/")

    return `${year}-${addZeros(month)}-${addZeros(day)}`
}

const addZeros = (str: string): string => {
    return parseInt(str) < 10 ? `0${str}` : str
}