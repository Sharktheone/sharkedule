import {createStyles} from "@mantine/core"


export const useColors = createStyles((theme) => ({
    loader: {
        "--loader-color-primary": theme.colors.blue[7] + "90",
        "--loader-color-secondary": theme.colors.blue[4] + "90",
    }
}))