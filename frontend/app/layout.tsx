/* eslint-disable @next/next/no-head-element */

import "./globals.scss";
import { Noto_Sans } from "@next/font/google";

const notoSans = Noto_Sans({
  weight: "900",
  subsets: ["latin"],
  variable: "--noto-font",
});

interface RootLayoutProps {
  children: React.ReactNode;
}

export default function RootLayout({ children }: RootLayoutProps) {
  return (
    <html lang="en" className={notoSans.variable}>
      <body>
        <main className="background">{children}</main>
      </body>
    </html>
  );
}
