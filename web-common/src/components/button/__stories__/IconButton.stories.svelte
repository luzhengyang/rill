<script lang="ts">
  import { Story, Template } from "@storybook/addon-svelte-csf";

  import IconButton from "@rilldata/web-common/components/button/IconButton.svelte";
  import SummaryStatistics from "@rilldata/web-common/components/icons/SummaryStatistics.svelte";

  export const meta = {
    title: 'Icon Button stories"',
    component: IconButton,
  };

  type ButtonProps = {
    disabled: boolean;
    rounded: boolean;
    size: number;
    bgGray: boolean;
    active: boolean;
  };

  const buttonProps: ButtonProps[] = [];

  for (const disabled of [true, false]) {
    for (const rounded of [true, false]) {
      for (const size of [20, 24, 28]) {
        for (const bgGray of [true, false]) {
          for (const active of [true, false]) {
            buttonProps.push({
              disabled,
              rounded,
              size,
              bgGray,
              active,
            });
          }
        }
      }
    }
  }
</script>

<Template>
  <table>
    <tr>
      <td />
      <td>disabled</td>
      <td>rounded</td>
      <td>size</td>
      <td>bgGray</td>
      <td>active</td>
    </tr>
    {#each buttonProps as props}
      <tr>
        <td
          ><IconButton {...props}
            ><SummaryStatistics />
            <svelte:fragment slot="tooltip-content">
              Show basic summary statistics
            </svelte:fragment></IconButton
          ></td
        >
        <td>{props.disabled}</td>
        <td>{props.rounded}</td>
        <td>{props.size}</td>
        <td>{props.bgGray}</td>
        <td>{props.active}</td>
      </tr>
    {/each}
  </table>
</Template>

<Story name="all icon button variations" />

<style>
  td {
    padding: 5px;
  }
  td:first-child {
    padding-right: 40px;
  }
</style>
