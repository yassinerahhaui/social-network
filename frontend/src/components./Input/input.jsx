export const InputWithLabel = ({
  id,
  type = "text",
  children,
}) => {
  return (
    <>
      <label htmlFor={id}>{children}</label>
      &nbsp;
      <input
        id={id}
        type={type}
      />
    </>
  );
};


export function InputField({ label, type, name, value, onChange, required }) {
  return (
    <div>
      <input
        type={type}
        name={name}
        placeholder={label}
        value={value}
        onChange={onChange}
        required={required}
      />
    </div>
  );
}