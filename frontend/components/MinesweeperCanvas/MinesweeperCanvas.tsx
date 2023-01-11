"use client";

// TODO: probably move this use client up to correct level

import { useEffect, useRef } from "react";

const SQUARE_SIZE = 50;

const FONT_SIZE = SQUARE_SIZE * 0.8;

const SQUARE_COLOUR = "#B9B9B9";

const HIDDEN_BORDER_WIDTH = 5;
const HIDDEN_COLOUR_HIGHLIGHT = "#FDFDFD";
const HIDDEN_COLOUR_SHADOW = "#767676";

const REVEALED_BORDER_COLOUR = "#777777";
const REVEALED_BORDER_WIDTH = 1;

const NUMBER_COLOURS = new Map<number, string>([
  [1, "#0000FF"],
  [2, "#008000"],
  [3, "#FF0000"],
  [4, "#000080"],
  [5, "#800000"],
  [6, "#008080"],
  [7, "#000000"],
  [8, "#808080"],
]);

type MinesweeperCanvasProps = {
  board: string[][];
};

const drawSquare = (ctx: CanvasRenderingContext2D, x: number, y: number) => {
  ctx.fillStyle = HIDDEN_COLOUR_HIGHLIGHT;
  ctx.beginPath();
  ctx.moveTo(x, y);
  ctx.lineTo(x + SQUARE_SIZE, y);
  ctx.lineTo(x, y + SQUARE_SIZE);
  ctx.fill();

  ctx.fillStyle = HIDDEN_COLOUR_SHADOW;
  ctx.beginPath();
  ctx.moveTo(x + SQUARE_SIZE, y + SQUARE_SIZE);
  ctx.lineTo(x + SQUARE_SIZE, y);
  ctx.lineTo(x, y + SQUARE_SIZE);
  ctx.fill();

  ctx.fillStyle = SQUARE_COLOUR;
  ctx.fillRect(
    x + HIDDEN_BORDER_WIDTH,
    y + HIDDEN_BORDER_WIDTH,
    SQUARE_SIZE - 2 * HIDDEN_BORDER_WIDTH,
    SQUARE_SIZE - 2 * HIDDEN_BORDER_WIDTH
  );
};

const drawEmptySquare = (
  ctx: CanvasRenderingContext2D,
  x: number,
  y: number
) => {
  // Render outer square border
  ctx.fillStyle = REVEALED_BORDER_COLOUR;
  ctx.fillRect(x, y, SQUARE_SIZE, SQUARE_SIZE);

  // Render inner square colour
  ctx.fillStyle = SQUARE_COLOUR;
  ctx.fillRect(
    x + REVEALED_BORDER_WIDTH,
    y + REVEALED_BORDER_WIDTH,
    SQUARE_SIZE - 2 * REVEALED_BORDER_WIDTH,
    SQUARE_SIZE - 2 * REVEALED_BORDER_WIDTH
  );
};

const drawNumber = (
  ctx: CanvasRenderingContext2D,
  value: number,
  x: number,
  y: number
) => {
  drawEmptySquare(ctx, x, y);

  ctx.fillStyle = NUMBER_COLOURS.get(value)!;
  ctx.textAlign = "center";
  ctx.font = `${FONT_SIZE}px Arial`;
  ctx.fillText(String(value), x + SQUARE_SIZE / 2, y + FONT_SIZE);
};

const drawBoard = (ctx: CanvasRenderingContext2D, board: string[][]) => {
  board.forEach((row, rowIndex) => {
    row.forEach((char, colIndex) => {
      const x = SQUARE_SIZE * colIndex;
      const y = SQUARE_SIZE * rowIndex;

      switch (char) {
        case "H":
          drawSquare(ctx, x, y);
          break;
        case "-":
          drawEmptySquare(ctx, x, y);
          break;
        default:
          if (!char.match(/[1-8]/)) {
            throw new Error(`Invalid character: ${char}`);
          }

          drawNumber(ctx, Number(char), x, y);
          break;
      }
    });
  });
};

const MinesweeperCanvas = ({ board }: MinesweeperCanvasProps) => {
  const canvasRef = useRef(null);

  useEffect(() => {
    if (!canvasRef.current) return;

    const canvas: HTMLCanvasElement = canvasRef.current;
    const context = canvas.getContext("2d");

    drawBoard(context!, board);
  }, []);

  return <canvas ref={canvasRef} width="250" height="250" />;
};

export default MinesweeperCanvas;
