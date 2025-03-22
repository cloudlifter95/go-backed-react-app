import React from "react";

interface Element {
  [key: string]: string | boolean;
}

interface ListProps {
  e_name: string;
  elements: Element[];
  headers: string[];
}

const ElementList: React.FC<ListProps> = ({ e_name, elements, headers }) => {
  return (
    <div className="p-4">
      <h1 className="text-xl font-bold mb-4"> {e_name} List</h1>
      <table>
        <thead>
          <tr>
            {headers.map((header, index) => (
              <th key={index}>{header}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {elements.map((e, rowIndex) => (
            <tr key={rowIndex}>
              {headers.map((header, colIndex) => (
                <td key={colIndex}>
                  {header === "Completed"
                    ? e[header]
                      ? "✅"
                      : "❌"
                    : e[header]}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default ElementList;
