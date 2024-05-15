# demo4_i18n

1. Create an empty message file for the language that you want to add (e.g. `translate.en.toml`).
2. Run `goi18n merge active.en.toml translate.en.toml` to populate `translate.en.toml` with the messages to be translated.

   ```toml
   # translate.en.toml
   [HelloPerson]
   hash = "sha1-5b49bfdad81fedaeefb224b0ffc2acc58b09cff5"
   other = "Hello {{.Name}}"
   ```

3. After `translate.en.toml` has been translated, rename it to `active.en.toml`.

   ```toml
   # active.en.toml
   [HelloPerson]
   hash = "sha1-5b49bfdad81fedaeefb224b0ffc2acc58b09cff5"
   other = "Hola {{.Name}}"
   ```

4. Load `active.en.toml` into your bundle.

   ```go
   bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
   bundle.LoadMessageFile("active.en.toml")
   ```
