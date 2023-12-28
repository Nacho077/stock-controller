import { useDispatch } from "react-redux";
import type { AppDispatch } from '../configs/redux/store'

const useAppDispatch: () => AppDispatch = useDispatch

export default useAppDispatch
