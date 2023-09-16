import {createStyles} from "@mantine/core"


export const useColors = createStyles((theme) => ({
    selector: {
        backgroundColor: theme.colors.dark[6],
        border: `1px solid ${
            theme.colorScheme === "dark" ? theme.colors.dark[5] : theme.colors.gray[2]
        }`,
        boxShadow: theme.shadows.sm,
    },
    color: {
        boxShadow: theme.shadows.sm,
    }
}))