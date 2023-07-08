import {createStyles} from "@mantine/core"


export const useColors = createStyles((theme) => ({
    description: {
        color: theme.colorScheme === "dark" ? theme.colors.dark[0] : theme.colors.gray[7],
        border: `1px solid ${
            theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
    }

}))