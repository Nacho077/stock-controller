import { TypedUseSelectorHook, useSelector } from "react-redux";
import { RootState } from "../configs/redux/store";

const useAppSelector: TypedUseSelectorHook<RootState> = useSelector

export default useAppSelector
